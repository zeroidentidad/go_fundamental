$(document).ready(function () {
    $("#form-registro").on("submit", function () {
        e.preventDefault();
        username = $('#username').val();

        $.ajax({
            type: 'POST',
            url: 'http://localhost:8000/validar',
            data: {
                "username": username
            },
            success: function(data) {
                result()
            }
        })
    })

    function result() {
        console.log("El servidor envio algo")
    }
})