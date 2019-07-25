//alert('Chidowan')
const cont = document.getElementById('contenido');
const btn = document.getElementById('btnAjax');

btn.addEventListener('click', e=>{
    const xhr = new XMLHttpRequest()
    xhr.open('GET','/data.html', true)

    // como usar lo solicitado
    xhr.addEventListener('load', e =>{
        cont.innerHTML=e.target.responseText
    })

    //Peticion
    xhr.send()
})