//Peticiones AJAX:

//Params: (metodo, url_solicitud, info_enviar)
function peticionAJAX(metodo, url, obj){
    // Con promesas
    return new Promise(function (resolver, rechazar){
        let xhr = new XMLHttpRequest();
        //con que metodo y adonde la peticion
        xhr.open(metodo, url, true);
        //encabezado de tipo de contenido a obtener
        xhr.setRequestHeader('Content-Type', 'application/json');

        if(sessionStorage.getItem('token')){
            xhr.setRequestHeader('Authorization', sessionStorage.getItem('token'));
        }

        xhr.addEventListener('load', e => {
            let self = e.target;
            let respuesta = {
                status: self.status,
                response: JSON.parse(self.response)
            };

            resolver(respuesta);
        });

        xhr.addEventListener('error', e => {
            let self = e.target;
            //console.log(self);

            rechazar(self);
        });

        xhr.send(obj);
    });
}

//Obtener objetos DOM:

function $(elemento) {
    return document.getElementById(elemento);
}