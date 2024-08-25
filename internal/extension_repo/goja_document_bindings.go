package extension_repo

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/dop251/goja"
	"strings"
)

func gojaBindDocument(vm *goja.Runtime) error {
	err := vm.Set("Doc", func(call goja.ConstructorCall) *goja.Object {
		gojaDoc := &GojaDoc{
			vm:            vm,
			doc:           nil,
			gojaSelection: nil,
		}
		obj := call.This

		if len(call.Arguments) != 1 {
			return goja.Undefined().ToObject(vm)
		}

		html := call.Arguments[0].String()

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			return goja.Undefined().ToObject(vm)
		}

		gojaDoc.doc = doc
		gojaDoc.gojaSelection = &GojaDocSelection{
			gojaDoc:   gojaDoc,
			selection: doc.Selection,
		}

		obj.Set("length", gojaDoc.gojaSelection.Length)
		obj.Set("html", gojaDoc.gojaSelection.Html)
		obj.Set("text", gojaDoc.gojaSelection.Text)
		obj.Set("attr", gojaDoc.gojaSelection.Attr)
		obj.Set("find", gojaDoc.gojaSelection.Find)
		obj.Set("children", gojaDoc.gojaSelection.Children)
		obj.Set("each", gojaDoc.gojaSelection.Each)
		obj.Set("text", gojaDoc.gojaSelection.Text)
		obj.Set("parent", gojaDoc.gojaSelection.Parent)
		obj.Set("parentsUntil", gojaDoc.gojaSelection.ParentsUntil)
		obj.Set("parents", gojaDoc.gojaSelection.Parents)
		obj.Set("end", gojaDoc.gojaSelection.End)
		obj.Set("closest", gojaDoc.gojaSelection.Closest)
		obj.Set("map", gojaDoc.gojaSelection.Map)
		obj.Set("first", gojaDoc.gojaSelection.First)
		obj.Set("last", gojaDoc.gojaSelection.Last)
		obj.Set("eq", gojaDoc.gojaSelection.Eq)
		obj.Set("contents", gojaDoc.gojaSelection.Contents)
		obj.Set("contentsFiltered", gojaDoc.gojaSelection.ContentsFiltered)
		obj.Set("filter", gojaDoc.gojaSelection.Filter)
		obj.Set("not", gojaDoc.gojaSelection.Not)
		obj.Set("is", gojaDoc.gojaSelection.Is)
		obj.Set("has", gojaDoc.gojaSelection.Has)
		obj.Set("next", gojaDoc.gojaSelection.Next)
		obj.Set("nextAll", gojaDoc.gojaSelection.NextAll)
		obj.Set("nextUntil", gojaDoc.gojaSelection.NextUntil)
		obj.Set("prev", gojaDoc.gojaSelection.Prev)
		obj.Set("prevAll", gojaDoc.gojaSelection.PrevAll)
		obj.Set("prevUntil", gojaDoc.gojaSelection.PrevUntil)
		obj.Set("siblings", gojaDoc.gojaSelection.Siblings)
		return obj
	})
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type GojaDoc struct {
	vm            *goja.Runtime
	doc           *goquery.Document
	gojaSelection *GojaDocSelection
}

type GojaDocSelection struct {
	gojaDoc   *GojaDoc
	selection *goquery.Selection
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Document
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func newGojaDocSelectionValue(d *GojaDoc, selection *goquery.Selection) goja.Value {
	gojaDocSelection := &GojaDocSelection{
		gojaDoc:   d,
		selection: selection,
	}

	obj := d.vm.NewObject()
	obj.Set("length", gojaDocSelection.Length)
	obj.Set("html", gojaDocSelection.Html)
	obj.Set("text", gojaDocSelection.Text)
	obj.Set("attr", gojaDocSelection.Attr)
	obj.Set("find", gojaDocSelection.Find)
	obj.Set("children", gojaDocSelection.Children)
	obj.Set("each", gojaDocSelection.Each)
	obj.Set("text", gojaDocSelection.Text)
	obj.Set("parent", gojaDocSelection.Parent)
	obj.Set("parentsUntil", gojaDocSelection.ParentsUntil)
	obj.Set("parents", gojaDocSelection.Parents)
	obj.Set("end", gojaDocSelection.End)
	obj.Set("closest", gojaDocSelection.Closest)
	obj.Set("map", gojaDocSelection.Map)
	obj.Set("first", gojaDocSelection.First)
	obj.Set("last", gojaDocSelection.Last)
	obj.Set("eq", gojaDocSelection.Eq)
	obj.Set("contents", gojaDocSelection.Contents)
	obj.Set("contentsFiltered", gojaDocSelection.ContentsFiltered)
	obj.Set("filter", gojaDocSelection.Filter)
	obj.Set("not", gojaDocSelection.Not)
	obj.Set("is", gojaDocSelection.Is)
	obj.Set("has", gojaDocSelection.Has)
	obj.Set("next", gojaDocSelection.Next)
	obj.Set("nextAll", gojaDocSelection.NextAll)
	obj.Set("nextUntil", gojaDocSelection.NextUntil)
	obj.Set("prev", gojaDocSelection.Prev)
	obj.Set("prevAll", gojaDocSelection.PrevAll)
	obj.Set("prevUntil", gojaDocSelection.PrevUntil)
	obj.Set("siblings", gojaDocSelection.Siblings)

	return obj
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Selection
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *GojaDocSelection) getFirstStringArg(call goja.FunctionCall) string {
	selectorStr, ok := call.Argument(0).Export().(string)
	if !ok {
		panic(s.gojaDoc.vm.NewTypeError("argument is not a string").ToString())
	}
	return selectorStr
}

func (s *GojaDocSelection) Length(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		return s.gojaDoc.vm.ToValue(0)
	}
	return s.gojaDoc.vm.ToValue(s.selection.Length())
}

// Find gets the descendants of each element in the current set of matched elements, filtered by a selector.
//
//	find(selector: string): DocSelection;
func (s *GojaDocSelection) Find(call goja.FunctionCall) (ret goja.Value) {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.Find(selectorStr))
}

func (s *GojaDocSelection) Html(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		return goja.Null()
	}
	htmlStr, err := s.selection.Html()
	if err != nil {
		return goja.Null()
	}
	return s.gojaDoc.vm.ToValue(htmlStr)
}

func (s *GojaDocSelection) Text(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		return s.gojaDoc.vm.ToValue("")
	}
	return s.gojaDoc.vm.ToValue(s.selection.Text())
}

// Attr gets the specified attribute's value for the first element in the Selection. To get the value for each element individually, use a
// looping construct such as Each or Map method.
//
//	attr(name: string): string | undefined;
func (s *GojaDocSelection) Attr(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	attr, found := s.selection.Attr(s.getFirstStringArg(call))
	if !found {
		return goja.Undefined()
	}
	return s.gojaDoc.vm.ToValue(attr)
}

// Parent gets the parent of each element in the Selection. It returns a new Selection object containing the matched elements.
//
//	parent(selector?: string): DocSelection;
func (s *GojaDocSelection) Parent(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.Parent())
	}

	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.ParentFiltered(selectorStr))
}

// Parents gets the ancestors of each element in the current Selection. It returns a new Selection object with the matched elements.
//
//	parents(selector?: string): DocSelection;
func (s *GojaDocSelection) Parents(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.Parents())
	}

	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.ParentsFiltered(selectorStr))
}

// ParentsUntil gets the ancestors of each element in the Selection, up to but not including the element matched by the selector. It returns a
// new Selection object containing the matched elements.
//
//	parentsUntil(selector?: string, until?: string): DocSelection;
func (s *GojaDocSelection) ParentsUntil(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	selectorStr := s.getFirstStringArg(call)
	if len(call.Arguments) < 2 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.ParentsUntil(selectorStr))
	}
	untilStr := call.Argument(1).String()
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.ParentsFilteredUntil(selectorStr, untilStr))
}

// End ends the most recent filtering operation in the current chain and returns the set of matched elements to its previous state.
//
//	end(): DocSelection;
func (s *GojaDocSelection) End(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.End())
}

// Closest gets the first element that matches the selector by testing the element itself and traversing up through its ancestors in the DOM tree.
//
//	closest(selector?: string): DocSelection;
func (s *GojaDocSelection) Closest(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	if len(call.Arguments) == 0 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.Closest(""))
	}

	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.Closest(selectorStr))
}

// Next gets the next sibling of each selected element, optionally filtered by a selector.
//
//	next(selector?: string): DocSelection;
func (s *GojaDocSelection) Next(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.Next())
	}

	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.NextFiltered(selectorStr))
}

// NextAll gets all following siblings of each element in the Selection, optionally filtered by a selector.
//
//	nextAll(selector?: string): DocSelection;
func (s *GojaDocSelection) NextAll(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.NextAll())
	}

	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.NextAllFiltered(selectorStr))
}

// NextUntil  gets all following siblings of each element up to but not including the element matched by the selector.
//
//	nextUntil(selector: string, until?: string): DocSelection;
func (s *GojaDocSelection) NextUntil(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	selectorStr := s.getFirstStringArg(call)
	if len(call.Arguments) < 2 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.NextUntil(selectorStr))
	}
	untilStr := call.Argument(1).String()
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.NextFilteredUntil(selectorStr, untilStr))
}

// Prev gets the previous sibling of each selected element optionally filtered by a selector.
//
//	prev(selector?: string): DocSelection;
func (s *GojaDocSelection) Prev(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.Prev())
	}

	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.PrevFiltered(selectorStr))
}

// PrevAll gets all preceding siblings of each element in the Selection, optionally filtered by a selector.
//
//	prevAll(selector?: string): DocSelection;
func (s *GojaDocSelection) PrevAll(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.PrevAll())
	}

	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.PrevAllFiltered(selectorStr))
}

// PrevUntil gets all preceding siblings of each element up to but not including the element matched by the selector.
//
//	prevUntil(selector: string, until?: string): DocSelection;
func (s *GojaDocSelection) PrevUntil(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	selectorStr := s.getFirstStringArg(call)
	if len(call.Arguments) < 2 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.PrevUntil(selectorStr))
	}
	untilStr := call.Argument(1).String()
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.PrevFilteredUntil(selectorStr, untilStr))
}

// Siblings gets the siblings of each element (excluding the element) in the set of matched elements, optionally filtered by a selector.
//
//	siblings(selector?: string): DocSelection;
func (s *GojaDocSelection) Siblings(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.Siblings())
	}

	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.SiblingsFiltered(selectorStr))
}

// Children gets the element children of each element in the set of matched elements.
//
//	children(selector?: string): DocSelection;
func (s *GojaDocSelection) Children(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.Children())
	}

	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.ChildrenFiltered(selectorStr))
}

// Contents gets the children of each element in the Selection, including text and comment nodes. It returns a new Selection object containing
// these elements.
//
//	contents(): DocSelection;
func (s *GojaDocSelection) Contents(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.Contents())
}

// ContentsFiltered gets the children of each element in the Selection, filtered by the specified selector. It returns a new Selection object
// containing these elements. Since selectors only act on Element nodes, this function is an alias to ChildrenFiltered unless the selector is
// empty, in which case it is an alias to Contents.
//
//	contentsFiltered(selector: string): DocSelection;
func (s *GojaDocSelection) ContentsFiltered(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.ContentsFiltered(selectorStr))
}

// Filter reduces the set of matched elements to those that match the selector string. It returns a new Selection object for this subset of
// matching elements.
//
//	filter(selector: string | (index: number, element: DocSelection) => boolean): DocSelection;
func (s *GojaDocSelection) Filter(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		panic(s.gojaDoc.vm.ToValue("missing argument"))
	}

	switch call.Argument(0).Export().(type) {
	case string:
		selectorStr := s.getFirstStringArg(call)
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.Filter(selectorStr))

	case func(call goja.FunctionCall) goja.Value:
		callback := call.Argument(0).Export().(func(call goja.FunctionCall) goja.Value)
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.FilterFunction(func(i int, selection *goquery.Selection) bool {
			ret, ok := callback(goja.FunctionCall{Arguments: []goja.Value{
				s.gojaDoc.vm.ToValue(i),
				newGojaDocSelectionValue(s.gojaDoc, selection),
			}}).Export().(bool)
			if !ok {
				panic(s.gojaDoc.vm.NewTypeError("callback did not return a boolean").ToString())
			}
			return ret
		}))
	default:
		panic(s.gojaDoc.vm.NewTypeError("argument is not a string or function").ToString())
	}
}

// Not removes elements from the Selection that match the selector string. It returns a new Selection object with the matching elements removed.
//
//	not(selector: string | (index: number, element: DocSelection) => boolean): DocSelection;
func (s *GojaDocSelection) Not(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		panic(s.gojaDoc.vm.ToValue("missing argument"))
	}

	switch call.Argument(0).Export().(type) {
	case string:
		selectorStr := s.getFirstStringArg(call)
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.Not(selectorStr))
	case func(call goja.FunctionCall) goja.Value:
		callback := call.Argument(0).Export().(func(call goja.FunctionCall) goja.Value)
		return newGojaDocSelectionValue(s.gojaDoc, s.selection.NotFunction(func(i int, selection *goquery.Selection) bool {
			ret, ok := callback(goja.FunctionCall{Arguments: []goja.Value{
				s.gojaDoc.vm.ToValue(i),
				newGojaDocSelectionValue(s.gojaDoc, selection),
			}}).Export().(bool)
			if !ok {
				panic(s.gojaDoc.vm.NewTypeError("callback did not return a boolean").ToString())
			}
			return ret
		}))
	default:
		panic(s.gojaDoc.vm.NewTypeError("argument is not a string or function").ToString())
	}
}

// Is checks the current matched set of elements against a selector and returns true if at least one of these elements matches.
//
//	is(selector: string | (index: number, element: DocSelection) => boolean): boolean;
func (s *GojaDocSelection) Is(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}

	if len(call.Arguments) == 0 {
		panic(s.gojaDoc.vm.ToValue("missing argument"))
	}

	switch call.Argument(0).Export().(type) {
	case string:
		selectorStr := s.getFirstStringArg(call)
		return s.gojaDoc.vm.ToValue(s.selection.Is(selectorStr))
	case func(call goja.FunctionCall) goja.Value:
		callback := call.Argument(0).Export().(func(call goja.FunctionCall) goja.Value)
		return s.gojaDoc.vm.ToValue(s.selection.IsFunction(func(i int, selection *goquery.Selection) bool {
			ret, ok := callback(goja.FunctionCall{Arguments: []goja.Value{
				s.gojaDoc.vm.ToValue(i),
				newGojaDocSelectionValue(s.gojaDoc, selection),
			}}).Export().(bool)
			if !ok {
				panic(s.gojaDoc.vm.NewTypeError("callback did not return a boolean").ToString())
			}
			return ret
		}))
	default:
		panic(s.gojaDoc.vm.NewTypeError("argument is not a string or function").ToString())
	}
}

// Has reduces the set of matched elements to those that have a descendant that matches the selector. It returns a new Selection object with the
// matching elements.
//
//	has(selector: string): DocSelection;
func (s *GojaDocSelection) Has(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	selectorStr := s.getFirstStringArg(call)
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.Has(selectorStr))
}

// Each iterates over a Selection object, executing a function for each matched element. It returns the current Selection object. The function f
// is called for each element in the selection with the index of the element in that selection starting at 0, and a *Selection that contains only
// that element.
//
//	each(callback: (index: number, element: DocSelection) => void): DocSelection;
func (s *GojaDocSelection) Each(call goja.FunctionCall) (ret goja.Value) {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	callback, ok := call.Argument(0).Export().(func(call goja.FunctionCall) goja.Value)
	if !ok {
		panic(s.gojaDoc.vm.NewTypeError("argument is not a function").ToString())
	}
	s.selection.Each(func(i int, selection *goquery.Selection) {
		callback(goja.FunctionCall{Arguments: []goja.Value{
			s.gojaDoc.vm.ToValue(i),
			newGojaDocSelectionValue(s.gojaDoc, selection),
		}})
	})
	return goja.Undefined()
}

// Map passes each element in the current matched set through a function, producing a slice of string holding the returned values. The function f
// is called for each element in the selection with the index of the element in that selection starting at 0, and a *Selection that contains only
// that element.
//
//	map(callback: (index: number, element: DocSelection) => DocSelection): DocSelection[];
func (s *GojaDocSelection) Map(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	callback, ok := call.Argument(0).Export().(func(call goja.FunctionCall) goja.Value)
	if !ok {
		panic(s.gojaDoc.vm.NewTypeError("argument is not a function").ToString())
	}
	var retStr []interface{}
	var retDocSelection map[string]interface{}
	s.selection.Each(func(i int, selection *goquery.Selection) {
		val := callback(goja.FunctionCall{Arguments: []goja.Value{
			s.gojaDoc.vm.ToValue(i),
			newGojaDocSelectionValue(s.gojaDoc, selection),
		}})

		if valExport, ok := val.Export().(map[string]interface{}); ok {
			retDocSelection = valExport
		}
		retStr = append(retStr, val.Export())

	})
	if len(retStr) > 0 {
		return s.gojaDoc.vm.ToValue(retStr)
	}
	return s.gojaDoc.vm.ToValue(retDocSelection)
}

// First reduces the set of matched elements to the first in the set. It returns a new Selection object, and an empty Selection object if the
// selection is empty.
//
//	first(): DocSelection;
func (s *GojaDocSelection) First(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.First())
}

// Last reduces the set of matched elements to the last in the set. It returns a new Selection object, and an empty Selection object if the
// selection is empty.
//
//	last(): DocSelection;
func (s *GojaDocSelection) Last(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.Last())
}

// Eq reduces the set of matched elements to the one at the specified index. If a negative index is given, it counts backwards starting at the
// end of the set. It returns a new Selection object, and an empty Selection object if the index is invalid.
//
//	eq(index: number): DocSelection;
func (s *GojaDocSelection) Eq(call goja.FunctionCall) goja.Value {
	if s.selection == nil {
		panic(s.gojaDoc.vm.ToValue("selection is nil"))
	}
	index, ok := call.Argument(0).Export().(int64)
	if !ok {
		panic(s.gojaDoc.vm.NewTypeError("argument is not a number").String())
	}
	return newGojaDocSelectionValue(s.gojaDoc, s.selection.Eq(int(index)))
}
