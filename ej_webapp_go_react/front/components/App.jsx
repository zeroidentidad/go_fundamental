import React, { Component } from 'react';
import SeccionCanales from './canales/SeccionCanales.jsx'
import SeccionUsuarios from './usuarios/SeccionUsuarios.jsx';
import SeccionMensajes from './mensajes/SeccionMensajes.jsx';

class App extends Component {
    constructor(props){
        super(props);
        this.state = {
            canales: [],
            canalActivo: {},
            usuarios: [],
            mensajes: [],
            connected: false
        };
    }
    //test websocket >
    componentDidMount(){
        let ws = this.ws = new WebSocket('ws://echo.websocket.org');
        ws.onmessage = this.message.bind(this);
        ws.onopen = this.open.bind(this);
        ws.onclose = this.close.bind(this);
    }
    // handlers websocket test
    message(e){
        const event = JSON.parse(e.data);
        if(event.nombre === 'canal agregado'){
            this.nuevoCanal(event.data)
        }
    }
    open(){
        this.setState({ connected: true });
    }
    close(){
        this.setState({ connected: false });
    }
    nuevoCanal(canal){
        let {canales}=this.state;
        canales.push(canal);
        this.setState({ canales });
    }
    //< test websocket
    agregarCanal(nombre){
        let {canales}=this.state;
        // canales.push({id: canales.length, nombre});
        // this.setState({canales});
        // enviar al servidor
        let msg = {
            nombre: 'canal agregado',
            data: {
                id: canales.length,
                nombre
            }
        }
        this.ws.send(JSON.stringify(msg));
    }
    setCanal(canalActivo){
        this.setState({canalActivo});
        // obtener canales de mensajes
    }
    setNombreUsuario(nombre){
        let {usuarios} = this.state;
        usuarios.push({id:usuarios.length, nombre});
        this.setState({usuarios});
        // enviar al servidor
    }
    agregarMensaje(body){
        let {mensajes, usuarios} = this.state;
        let createdAt = new Date;
        let autor = usuarios.length > 0 ? usuarios[0].nombre: 'anonimo';
        mensajes.push({id:mensajes.length, body, createdAt, autor});
        this.setState({mensajes});
        // enviar al servidor
    }
    render() {
        return (
            <div className='app'>
                <div className='nav'>
                    <SeccionCanales
                        /*canales={this.state.canales}*/
                        {...this.state}
                        agregarCanal={this.agregarCanal.bind(this)}
                        setCanal={this.setCanal.bind(this)}
                    />
                    <SeccionUsuarios 
                        {...this.state}
                        setNombreUsuario={this.setNombreUsuario.bind(this)}
                    />
                </div>
                <SeccionMensajes
                    {...this.state}
                    agregarMensaje={this.agregarMensaje.bind(this)}
                />
            </div>
        )
    }
}

export default App;