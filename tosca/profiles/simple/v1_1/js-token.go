// This file was auto-generated from YAML files

package v1_1

func init() {
	Profile["/tosca/simple/1.1/js/token.js"] = `

function evaluate(v, separators, index) {
	if (arguments.length !== 3)
		throw 'must have 3 arguments';
	if (v.$string !== undefined)
		v = v.$string;
	s = v.split(new RegExp('[' + escape(separators) + ']'));
	return s[index];
}

function escape(s) {
	return s.replace(/[\-\[\]\/\{\}\(\)\*\+\?\.\\\^\$\|]/g, '\\$&');
}
`
}
