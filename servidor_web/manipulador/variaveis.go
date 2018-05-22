package manipulador

import "html/template"

//Modelos armazenam os modelos de pagina que serao executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
