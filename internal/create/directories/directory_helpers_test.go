package directories

import "os"

var tests = []struct {
	input             string
	errorShouldHappen bool
	path              string
}{
	{
		input: `.
   |-- src
      |- index.js
      |-- components
         |- header.js
   |-- test
      |- <<test-name>>.js
      |-- fixtures
         |- test.fixture.js
`,
		errorShouldHappen: false,
		path: os.TempDir(),
	},
	{
		input: `express
   |-- src
      |- index.js
      |-- components
         |- header.js
   |-- test
      |- <<test-name>>.js
      |-- fixtures
         |- test.fixture.js
		`,
		errorShouldHappen: false,
		path: os.TempDir(),
	},
	{
		input: `express`,
		errorShouldHappen: false,
		path: os.TempDir(),
	},
	{
		input: `projectRoot
   |-- src
     |- index.js
secondRoot
		`,
		errorShouldHappen: true,
		path: os.TempDir(),
	},
	{
		input: `projectRoot
secondRoot
		`,
		errorShouldHappen: true,
		path: os.TempDir(),
	},
	{
		input: `projectRoot
   |-- src
	    |- index.js
		`,
		errorShouldHappen: true,
		path: os.TempDir(),
	},
}