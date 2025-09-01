package main
import (
	"net/http"
)

func (app *application) healthcheckHandler( w http.ResponseWriter,r *http.Request){
	// fmt.Fprintln(w,"status: available")
	// fmt.Fprintf(w,"enviroment: %s\n",app.config.env)
	// fmt.Fprintf(w,"version:%s\n",version)
	data := map[string]string{
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}
	// js := `{"status" : "available","environment": %q, "version" : %q}`
	// js = fmt.Sprintf(js, app.config.env, version)
	// w.Header().Set("Content-Type","application/json")
	
	// w.Write([]byte(js))
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil{
		app.logger.Println(err)
		http.Error(w,"the server encountered a problem and could not process your request",http.StatusInternalServerError)
		return
	}
// 	js = append(js, '\n')
// 	w.Header().Set("Content-Type", "application/json")

//  w.Write(js)
}
