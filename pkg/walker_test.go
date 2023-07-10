package pkg

import (
	"bytes"
	"context"
	"io"
	"os"
	"regexp"
	"testing"

	"github.com/vektra/mockery/v2/pkg/config"
	"gitlab.com/incubus8/gotest/assert"
)

type GatheringVisitor struct {
	Interfaces []*Interface
}

func (v *GatheringVisitor) VisitWalk(ctx context.Context, iface *Interface) error {
	v.Interfaces = append(v.Interfaces, iface)
	return nil
}

func NewGatheringVisitor() *GatheringVisitor {
	return &GatheringVisitor{
		Interfaces: make([]*Interface, 0, 1024),
	}
}

type BufferedProvider struct {
	buf *bytes.Buffer
}

func NewBufferedProvider() *BufferedProvider {
	return &BufferedProvider{
		buf: new(bytes.Buffer),
	}
}

func (bp *BufferedProvider) String() string {
	return bp.buf.String()
}

func (bp *BufferedProvider) GetWriter(context.Context, *Interface) (io.Writer, error, Cleanup) {
	return bp.buf, nil, func() error { return nil }
}

func TestWalkerHere(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping recursive walker test")
	}

	wd, err := os.Getwd()
	assert.NoError(t, err)
	w := Walker{
		BaseDir:   wd,
		Recursive: true,
		LimitOne:  false,
		Filter:    regexp.MustCompile(".*"),
	}

	gv := NewGatheringVisitor()

	w.Walk(context.Background(), gv)

	assert.True(t, len(gv.Interfaces) > 10)
	first := gv.Interfaces[0]
	assert.Equal(t, "A", first.Name)
	assert.Equal(t, getFixturePath("struct_value.go"), first.FileName)
	assert.Equal(t, "github.com/vektra/mockery/v2/pkg/fixtures", first.QualifiedName)
}

func TestWalkerRegexp(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping recursive walker test")
	}

	wd, err := os.Getwd()
	assert.NoError(t, err)
	w := Walker{
		BaseDir:   wd,
		Recursive: true,
		LimitOne:  false,
		Filter:    regexp.MustCompile(".*AsyncProducer*."),
	}

	gv := NewGatheringVisitor()

	w.Walk(context.Background(), gv)

	assert.True(t, len(gv.Interfaces) >= 1)
	first := gv.Interfaces[0]
	assert.Equal(t, "AsyncProducer", first.Name)
	assert.Equal(t, getFixturePath("async.go"), first.FileName)
	assert.Equal(t, "github.com/vektra/mockery/v2/pkg/fixtures", first.QualifiedName)
}

func TestPackagePrefix(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping recursive walker test")
	}

	wd, err := os.Getwd()
	assert.NoError(t, err)

	w := Walker{
		BaseDir:   wd,
		Recursive: true,
		LimitOne:  false,
		Filter:    regexp.MustCompile(".*AsyncProducer*."),
	}

	bufferedProvider := NewBufferedProvider()
	visitor := NewGeneratorVisitor(GeneratorVisitorConfig{
		InPackage:         false,
		PackageName:       "mocks",
		PackageNamePrefix: "prefix_test_",
	}, bufferedProvider, false)

	w.Walk(context.Background(), visitor)
	assert.Regexp(t, regexp.MustCompile("package prefix_test_test"), bufferedProvider.String())
}

func TestWalkerExclude(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping recursive walker test")
	}

	wd, err := os.Getwd()
	assert.NoError(t, err)

	w := Walker{
		BaseDir:   wd,
		Recursive: true,
		LimitOne:  false,
		Config: config.Config{
			Exclude: []string{
				getFixturePath("requester"),
				getFixturePath("generic.go"),
			},
		},
		Filter: regexp.MustCompile(".*"),
	}

	gv := NewGatheringVisitor()

	w.Walk(context.Background(), gv)
	for _, iface := range gv.Interfaces {
		assert.NotContains(t, iface.Name, "Requester",
			"Interface %s should have been excluded but found in file: %s", iface.Name, iface.FileName)
	}
}
