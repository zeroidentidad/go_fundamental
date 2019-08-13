let ws = null, theChart = null
const dataChart = [5, 7, 9]

const setSystemMessage = data => {
    systemMessage.textContent = data
}

const login = async() => {
    const user = {
        name: usrName.value,
        password: password.value
    }

    const header = new Headers()
    header.append('Content-Type', 'application/json')
    let data = {}
    const response = await fetch('/login', { method: 'POST', headers: header, body: JSON.stringify(user) })
    switch (response.status) {
        case 200:
           data = await response.json()
            connectWS(data)
            loadChart()
            setSystemMessage(`Conectado OK. Code: ${response.status}`)
            break;
        case 401:
            setSystemMessage(`Usuario o contraseña no valido. Code: ${response.status}`)
            break;            
        default:
            setSystemMessage(`Estado no esperado. Code: ${response.status}`);
            break;
    }
}

btnLogin.addEventListener('click', e => {
    e.preventDefault()
    login()
})

const connectWS = data => {
    ws = new WebSocket(`ws://localhost:9000/ws?uname=${usrName.value}&token=${data.token}`)
    ws.onopen = e => {
        setSystemMessage('Conectado a websocket correctamente.')
    }
    ws.onerror = e => {
        setSystemMessage(`Ocurrio un error. Desc: ${e}`)
    }
    ws.onmessage = e => {
        const data = JSON.parse(e.data)
        switch (data.type) {
            case 'message':
                content.insertAdjacentHTML('beforeend',`<div>De: ${data.data_response.name}, Mensaje: ${data.data_response.message}</div>`)
                break;
            case 'sale':
                dataChart[data.data_sale.product] += data.data_sale.quantity
                theChart.update()
                break;
            case 'pong':
                setSystemMessage('Reconexion al WS.')
                console.log('Reconexion a WS')
                break;                        
            default:
                setSystemMessage('Mensaje desconocido recibido')
                break;
        }
    }

    setInterval(() => {
        ws.send(JSON.stringify({type: 'ping'}))
    }, 60000);
}

btnSend.addEventListener('click', e => {
    e.preventDefault()
    let data={};
    if (e.data === null || e.data ===''){
        setSystemMessage('Mensaje vacio.')        
    }else{
        data = {
            type: 'message',
            message: textMsg.value
        }
        ws.send(JSON.stringify(data))
    }
})

btnSale.addEventListener('click', e => {
    e.preventDefault()
    const data = {
            type: 'sale',
            product: parseInt(product.value, 10),
            quantity: parseInt(quantity.value, 10)
    }
    if(data.quantity>0){
        ws.send(JSON.stringify(data)) 
    }else{
        setSystemMessage('Sin cantidad.') 
    }   
})

const loadChart = () => {
    const ctx = myChart.getContext('2d')
    myChart.width = 400
    myChart.height = 400
    theChart  = new Chart(ctx,{
        type: 'bar',
        fill: false,
        data:{
            labels: ["Caguamita", "Caguama", "Caguamon"],
            datasets:[{
                label: "Ventas",
                data: dataChart,
                backgroundColor: [
                    'rgba(255, 99, 132, 0.5)',
                    'rgba(54, 162, 235, 0.5)',
                    'rgba(255, 206, 86, 0.5)'
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1.5)',
                    'rgba(54, 162, 235, 1.5)',
                    'rgba(255, 206, 86, 1.5)'
                ],
                borderWidth: 1
            }]
        },
        options:{
            hoverMode: 'index',
            stacked: false,
            responsive: true,
            scales:{
                yAxes: [{
                    ticks:{
                        beginAtZero: true
                    }
                }]
            }
        }
    })
}

/*window.randomScalingFactor = function () {
    return (Math.random() > 0.5 ? 1.0 : -1.0) * Math.random() * 100;
};*/