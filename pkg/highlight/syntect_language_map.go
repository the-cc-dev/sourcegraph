package highlight

// SyntectLanguageMap is a map that maps language identifiers that may be provided at the end of markdown code fences
// to valid syntect file extensions.
var SyntectLanguageMap = map[string]string{"txt": "txt",
	"asa":                  "asa",
	"asp":                  "asp",
	"actionscript":         "as",
	"as":                   "as",
	"applescript":          "applescript",
	"script editor":        "script editor",
	"batchfile":            "bat",
	"bat":                  "bat",
	"cmd":                  "cmd",
	"build":                "build",
	"c#":                   "cs",
	"cs":                   "cs",
	"csx":                  "csx",
	"cpp":                  "cpp",
	"cc":                   "cc",
	"cp":                   "cp",
	"cxx":                  "cxx",
	"c++":                  "c++",
	"c":                    "c",
	"h":                    "h",
	"hh":                   "hh",
	"hpp":                  "hpp",
	"hxx":                  "hxx",
	"h++":                  "h++",
	"inl":                  "inl",
	"ipp":                  "ipp",
	"cmakecache.txt":       "cmakecache.txt",
	"cmakelists.txt":       "cmakelists.txt",
	"cmake":                "cmake",
	"css":                  "css",
	"css.erb":              "css.erb",
	"css.liquid":           "css.liquid",
	"capnp":                "capnp",
	"cg":                   "cg",
	"clojure":              "clj",
	"clj":                  "clj",
	"cljs":                 "cljs",
	"cljc":                 "cljc",
	"cljx":                 "cljx",
	"crontab":              "crontab",
	"d":                    "d",
	"di":                   "di",
	"diff":                 "diff",
	"patch":                "patch",
	"dockerfile":           "dockerfile",
	"erlang":               "erl",
	"erl":                  "erl",
	"hrl":                  "hrl",
	"emakefile":            "emakefile",
	"yaws":                 "yaws",
	"fsharp":               "fs",
	"forth":                "frt",
	"frt":                  "frt",
	"essl":                 "essl",
	"f.essl":               "f.essl",
	"v.essl":               "v.essl",
	"_v.essl":              "_v.essl",
	"_f.essl":              "_f.essl",
	"_vs.essl":             "_vs.essl",
	"_fs.essl":             "_fs.essl",
	"vs":                   "vs",
	"fs":                   "fs",
	"gs":                   "gs",
	"vsh":                  "vsh",
	"fsh":                  "fsh",
	"gsh":                  "gsh",
	"vshader":              "vshader",
	"fshader":              "fshader",
	"gshader":              "gshader",
	"vert":                 "vert",
	"frag":                 "frag",
	"geom":                 "geom",
	"tesc":                 "tesc",
	"tese":                 "tese",
	"comp":                 "comp",
	"glsl":                 "glsl",
	"attributes":           "attributes",
	"gitattributes":        "gitattributes",
	".gitattributes":       ".gitattributes",
	"commit_editmsg":       "commit_editmsg",
	"merge_msg":            "merge_msg",
	"tag_editmsg":          "tag_editmsg",
	"gitconfig":            "gitconfig",
	".gitconfig":           ".gitconfig",
	".gitmodules":          ".gitmodules",
	"exclude":              "exclude",
	"gitignore":            "gitignore",
	".gitignore":           ".gitignore",
	"gitlink":              ".git",
	".git":                 ".git",
	"gitlog":               "gitlog",
	"git-rebase-todo":      "git-rebase-todo",
	"go":                   "go",
	"graphviz-dot":         "dot",
	"dot":                  "dot",
	"gv":                   "gv",
	"groovy":               "groovy",
	"gvy":                  "gvy",
	"gradle":               "gradle",
	"fx":                   "fx",
	"fxh":                  "fxh",
	"hlsl":                 "hlsl",
	"hlsli":                "hlsli",
	"usf":                  "usf",
	"html":                 "html",
	"htm":                  "htm",
	"shtml":                "shtml",
	"xhtml":                "xhtml",
	"tmpl":                 "tmpl",
	"tpl":                  "tpl",
	"haskell":              "hs",
	"hs":                   "hs",
	"literatehaskell":      "lhs",
	"lhs":                  "lhs",
	"ini":                  "ini",
	"inf":                  "inf",
	"reg":                  "reg",
	"lng":                  "lng",
	"cfg":                  "cfg",
	"url":                  "url",
	".editorconfig":        ".editorconfig",
	"jsp":                  "jsp",
	"java":                 "java",
	"bsh":                  "bsh",
	"javaproperties":       "javaproperties",
	"properties":           "properties",
	"json":                 "json",
	"sublime-settings":     "sublime-settings",
	"sublime-menu":         "sublime-menu",
	"sublime-keymap":       "sublime-keymap",
	"sublime-mousemap":     "sublime-mousemap",
	"sublime-theme":        "sublime-theme",
	"sublime-build":        "sublime-build",
	"sublime-project":      "sublime-project",
	"sublime-completions":  "sublime-completions",
	"sublime-commands":     "sublime-commands",
	"sublime-macro":        "sublime-macro",
	"sublime-color-scheme": "sublime-color-scheme",
	"javascript":           "js",
	"js":                   "js",
	"jsx":                  "jsx",
	"babel":                "babel",
	"es6":                  "es6",
	"less":                 "less",
	"BibTeX":               "bib",
	"bib":                  "bib",
	"LaTeX":                "ltx",
	"latex":                "ltx",
	"tex":                  "tex",
	"ltx":                  "ltx",
	"TeX":                  "sty",
	"sty":                  "sty",
	"cls":                  "cls",
	"lisp":                 "lisp",
	"cl":                   "cl",
	"clisp":                "clisp",
	"l":                    "l",
	"mud":                  "mud",
	"el":                   "el",
	"scm":                  "scm",
	"ss":                   "ss",
	"lsp":                  "lsp",
	"fasl":                 "fasl",
	"lua":                  "lua",
	"proj":                 "proj",
	"targets":              "targets",
	"msbuild":              "msbuild",
	"csproj":               "csproj",
	"vbproj":               "vbproj",
	"fsproj":               "fsproj",
	"vcxproj":              "vcxproj",
	"make":                 "make",
	"gnumakefile":          "gnumakefile",
	"makefile":             "makefile",
	"ocamlmakefile":        "ocamlmakefile",
	"mak":                  "mak",
	"mk":                   "mk",
	"man":                  "man",
	"md":                   "md",
	"mdown":                "mdown",
	"markdown":             "markdown",
	"markdn":               "markdn",
	"matlab":               "matlab",
	"pom.xml":              "pom.xml",
	"mediawiki":            "mediawiki",
	"wikipedia":            "wikipedia",
	"wiki":                 "wiki",
	"ninja":                "ninja",
	"ocaml":                "ml",
	"ml":                   "ml",
	"mli":                  "mli",
	"ocamllex":             "mll",
	"mll":                  "mll",
	"ocamlyacc":            "mly",
	"mly":                  "mly",
	"objective-c++":        "mm",
	"mm":                   "mm",
	"objective-c":          "m",
	"m":                    "m",
	"php":                  "php",
	"php3":                 "php3",
	"php4":                 "php4",
	"php5":                 "php5",
	"php7":                 "php7",
	"phps":                 "phps",
	"phpt":                 "phpt",
	"phtml":                "phtml",
	"pascal":               "pas",
	"pas":                  "pas",
	"p":                    "p",
	"dpr":                  "dpr",
	"spec":                 "spec",
	"client":               "client",
	"perl":                 "pl",
	"pm":                   "pm",
	"pod":                  "pod",
	"t":                    "t",
	"pl":                   "pl",
	"postscript":           "ps",
	"ps":                   "ps",
	"eps":                  "eps",
	"powershell":           "ps1",
	"ps1":                  "ps1",
	"psm1":                 "psm1",
	"psd1":                 "psd1",
	"python":               "py",
	"py":                   "py",
	"py3":                  "py3",
	"pyw":                  "pyw",
	"pyi":                  "pyi",
	"pyx":                  "pyx",
	"pyx.in":               "pyx.in",
	"pxd":                  "pxd",
	"pxd.in":               "pxd.in",
	"pxi":                  "pxi",
	"pxi.in":               "pxi.in",
	"rpy":                  "rpy",
	"cpy":                  "cpy",
	"sconstruct":           "sconstruct",
	"sconscript":           "sconscript",
	"gyp":                  "gyp",
	"gypi":                 "gypi",
	"snakefile":            "snakefile",
	"wscript":              "wscript",
	"r":                    "r",
	"s":                    "s",
	"rprofile":             "rprofile",
	"rd":                   "rd",
	"rails":                "rails",
	"rhtml":                "rhtml",
	"erb":                  "erb",
	"html.erb":             "html.erb",
	"js.erb":               "js.erb",
	"haml":                 "haml",
	"rubyonrails":          "rxml",
	"rxml":                 "rxml",
	"builder":              "builder",
	"erbsql":               "erbsql",
	"sql.erb":              "sql.erb",
	"re":                   "re",
	"restructuredtext":     "rst",
	"rst":                  "rst",
	"rest":                 "rest",
	"ruby":                 "rb",
	"rb":                   "rb",
	"appfile":              "appfile",
	"appraisals":           "appraisals",
	"berksfile":            "berksfile",
	"brewfile":             "brewfile",
	"capfile":              "capfile",
	"cgi":                  "cgi",
	"cheffile":             "cheffile",
	"config.ru":            "config.ru",
	"deliverfile":          "deliverfile",
	"fastfile":             "fastfile",
	"fcgi":                 "fcgi",
	"gemfile":              "gemfile",
	"gemspec":              "gemspec",
	"guardfile":            "guardfile",
	"irbrc":                "irbrc",
	"jbuilder":             "jbuilder",
	"podspec":              "podspec",
	"prawn":                "prawn",
	"rabl":                 "rabl",
	"rake":                 "rake",
	"rakefile":             "rakefile",
	"rantfile":             "rantfile",
	"rbx":                  "rbx",
	"rjs":                  "rjs",
	"ruby.rail":            "ruby.rail",
	"scanfile":             "scanfile",
	"simplecov":            "simplecov",
	"snapfile":             "snapfile",
	"thor":                 "thor",
	"thorfile":             "thorfile",
	"vagrantfile":          "vagrantfile",
	"rs":                   "rs",
	"sass":                 "sass",
	"scss":                 "scss",
	"sql":                  "sql",
	"ddl":                  "ddl",
	"dml":                  "dml",
	"scala":                "scala",
	"sbt":                  "sbt",
	"sh":                   "sh",
	"bash":                 "bash",
	"zsh":                  "zsh",
	"fish":                 "fish",
	".bash_aliases":        ".bash_aliases",
	".bash_completions":    ".bash_completions",
	".bash_functions":      ".bash_functions",
	".bash_login":          ".bash_login",
	".bash_logout":         ".bash_logout",
	".bash_profile":        ".bash_profile",
	".bash_variables":      ".bash_variables",
	".bashrc":              ".bashrc",
	".profile":             ".profile",
	".textmate_init":       ".textmate_init",
	"smalltalk":            "st",
	"st":                   "st",
	"swift":                "swift",
	"adp":                  "adp",
	"tcl":                  "tcl",
	"toml":                 "toml",
	"tml":                  "tml",
	"lock":                 "lock",
	"textile":              "textile",
	"thrift":               "thrift",
	"typescript":           "ts",
	"ts":                   "ts",
	"typescriptreact":      "tsx",
	"tsx":                  "tsx",
	"xml":                  "xml",
	"xsd":                  "xsd",
	"xslt":                 "xslt",
	"tld":                  "tld",
	"dtml":                 "dtml",
	"rss":                  "rss",
	"opml":                 "opml",
	"svg":                  "svg",
	"yaml":                 "yaml",
	"yml":                  "yml",
	"sublime-syntax":       "sublime-syntax"}
