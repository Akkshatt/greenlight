package main
import (
	"net/http"
)

func (app *application) healthcheckHandler( w http.ResponseWriter,r *http.Request){
	// fmt.Fprintln(w,"status: available")
	// fmt.Fprintf(w,"enviroment: %s\n",app.config.env)
	// fmt.Fprintf(w,"version:%s\n",version)
	env := envelope{
"status": "available",
"system_info": map[string]string{
"environment": app.config.env,
"version": version,
},
}

	err := app.writeJSON(w, http.StatusOK, env, nil)
if err != nil {
// Use the new serverErrorResponse() helper.
app.serverErrorResponse(w, r, err)
}
}
