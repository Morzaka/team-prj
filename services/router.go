package services

import(
        "net/http"
        "github.com/go-zoo/bone"
)

func startPage() {
        mux := bone.New().Prefix("/api")
        sub := bone.New()
        mux.SubRoute("/v1", sub)
         // GetFunc, PostFunc etc ... takes http.HandlerFunc
        sub.GetFunc("/gogroup/:id", handlerFunc)
        http.ListenAndServe(":8000", mux)
}

func handlerFunc(w http.ResponseWriter, r *http.Request){
        val := bone.GetValue(r, "id")
        w.Write([]byte("Hello golang group " + val))
}
