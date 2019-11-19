package serializer

import (
	"github.com/flywave/go-osg"
	"github.com/flywave/go-osg/model"
)

func getShadeMode(obj interface{}) interface{} {
	return &obj.(*model.ShadeModel).Mode
}

func setShadeMode(obj interface{}, val interface{}) {
	obj.(*model.ShadeModel).Mode = val.(int)
}

func init() {
	fn := func() interface{} {
		sm := model.NewShadeModel()
		return &sm
	}
	wrap := osg.NewObjectWrapper2("ShadeModel", "flywave::osg::shademodel", fn, "osg::Object osg::StateAttribute osg::ShadeModel")
	ser := osg.NewEnumSerializer("Mode", getShadeMode, setShadeMode)
	ser.Add("FLAT", model.FLAT)
	ser.Add("SMOOTH", model.SMOOTH)
	wrap.AddSerializer(&ser, osg.RWENUM)
	osg.GetObjectWrapperManager().AddWrap(&wrap)
}