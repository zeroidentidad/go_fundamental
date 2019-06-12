
let formLogin = $('formLogin'),
    email = $('email'),
    password = $('password'),
    btnLogin = $('btnLogin'),
    mensajeLogin = $('mensaje-login');

formLogin.addEventListener('submit', e=>{
    e.preventDefault();
    let obj = {
        email: email.value,
        password: password.value   
    }

    peticionAJAX(formLogin.method, formLogin.action, JSON.stringify(obj))
    .then(respuesta =>{
        if(respuesta.status === 200){
            //console.log('Ingreso correcto:');
            mensajeLogin.textContent = 'Ingreso OK';

            sessionStorage.setItem('token', respuesta.response.token);

            console.log(respuesta.response);
        }else{
            mensajeLogin.textContent = respuesta.response.message
            console.log(respuesta.response);
        }
    })
    .catch(error => {
        console.log(error);
    });

});