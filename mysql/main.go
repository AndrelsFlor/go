package main

import (
	"flag"
	"encoding/json"
	"github.com/gorilla/mux"
	//"fmt"
	// "database/sql"  // Pacote Database SQL para realizar query
	//"log"           // Mostra informações no console
	"net/http"      // Gerencia URLs e Servidor Web
	//"text/template" // Gerencia templates

	//"github.com/go-sql-driver/mysql" // Driver Mysql para Go
	//"github.com/julienschmidt/httprouter"
)

var(
	addr = flag.String("addr",":8080","endereço do serviço http")
	data  map[string]string
)

type Names struct{
	id int
	name string
	email string
	
}

func main(){
	flag.Parse()

	data = map[string]string{}

	r := mux.NewRouter()
    r.HandleFunc("/cada/usuario", cadUsuario)
    http.ListenAndServe(*addr, r)
	
}

// func dbConn()(db *sql.Db){
// 	dbDriver := "mysql" // Driver do banco de dados
// 	dbUser := "root"        // Usuário
// 	dbPass := ""        // Senha
// 	dbName := "go"        // Nome do banco


// 	// Realiza a conexão com banco de dados:
// 	// sql.Open("DRIVER", "Usuario:Senha/@BancoDeDados)
// 	// A variavel `db` é utilizada junto com pacote `database/sql`
// 	// para a montagem de Query.
// 	// A variavel `err` é utilizada no tratamento de erros
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }


func cadUsuario(w http.ResponseWriter,r *http.Request){
	
	 
	data["nome"] 	= r.FormValue("nome");
	data["email"] 	= r.FormValue("email");
	data["senha"]   = r.FormValue("senha");

	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(data)

}

