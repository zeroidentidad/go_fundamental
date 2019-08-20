$(document).ready(function () {
    $("#form-registro").on("submit", function (e) {
        e.preventDefault();
        username = $('#username').val();

        $.ajax({
            type: 'POST',
            url: 'http://localhost:8000/validar',
            data: {
                "username": username
            },
            success: function(data) {
                result(data)
            }
        })
    })

    function result(data) {
        //console.log("El servidor envio algo, Data:", JSON.parse(data))
        obj = JSON.parse(data)

        if(obj.valid===true){
            crearConexion()
        }else{
            console.log("Reintentando :v")
        }
    }

    function crearConexion() {
        console.log("conectado")
    }
})