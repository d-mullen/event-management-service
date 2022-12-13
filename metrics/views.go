package metrics

import (
	"go.opencensus.io/stats/view"
)

var AllViews = []*view.View{
	// set status
	SetStatusTimeMsView,
	SetStatusCountView,
	// annotate
	AnnotateTimeMsView,
	AnnotateCountView,
	DeleteAnnotationCountView,
}
