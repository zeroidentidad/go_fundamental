
const persona = {
    nombre: document.getElementById('nombre').value,
    apellido: document.getElementById('apellido').value,
    edad: parseInt(document.getElementById('edad').value, 10),
    activo: document.getElementById('activo').checked
}

const misHeaders = new Headers();
misHeaders.append('Content-type', 'application/json')
misHeaders.append('Authorization','Bearer dfdf4rifjdfdfdf6565')

const config = {
    method: 'POST',
    headers: misHeaders,
    body: JSON.encode(persona)
}

fetch('/api', config)
.then(res=>res.json())
.then(res=>dibujar(res.data))