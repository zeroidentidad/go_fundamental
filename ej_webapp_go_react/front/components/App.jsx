import React, { Component } from 'react';
import SeccionCanales from './canales/SeccionCanales.jsx'
import SeccionUsuarios from './usuarios/SeccionUsuarios.jsx';
import SeccionMensajes from './mensajes/SeccionMensajes.jsx';
import Socket from '../socket.js';

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
        //let ws = this.ws = new WebSocket('ws://echo.websocket.org');
        let socket = this.socket = new Socket();
        socket.on('connect', this.onConnect.bind(this));
        socket.on('disconnect', this.onDisconnect.bind(this));
        socket.on('agregar canal', this.onAgregarCanal.bind(this));
        socket.on('agregar usuario', this.onAgregarUsuario.bind(this));
        socket.on('editar usuario', this.onEditarUsuario.bind(this));
        socket.on('remover usuario', this.onRemoverUsuario.bind(this));
        socket.on('agregar mensaje', this.onAgregarMensaje.bind(this));
    }
    onConnect(){
        this.setState({ connected: true });
        this.socket.emit('subscribir canal');
        this.socket.emit('subscribir usuario');
    }
    onDisconnect() {
        this.setState({ connected: false });
    }
    onAgregarUsuario(usuario){
        let { usuarios } = this.state;
        usuarios.push(usuario);
        this.setState({ usuarios });
    }
    onEditarUsuario(editarUsuario){
        let { usuarios } = this.state;
        usuarios = usuarios.map(usuario =>{
            if(editarUsuario.id===usuario.id){
                return editarUsuario;
            }
            return usuario
        });
        this.setState({ usuarios });
    }
    onRemoverUsuario(removerUsuario){
        let {usuarios}=this.state;
        usuarios = usuarios.filter(usuario=>{
            return usuario.id !== removerUsuario.id;
        });
        this.setState({ usuarios });
    }
    onAgregarMensaje(mensaje){
        let {mensajes} = this.state;
        mensajes.push(mensaje);
        this.setState({ mensajes });
    }    
    //nuevoCanal(canal){
    onAgregarCanal(canal){
        let {canales}=this.state;
        canales.push(canal);
        this.setState({ canales });
    }
    //< test websocket
    agregarCanal(nombre){
        //let {canales}=this.state;
        // canales.push({id: canales.length, nombre});
        // this.setState({canales});
        // enviar al servidor
        this.socket.emit('agregar canal',{nombre});
        /*let msg = {
            nombre: 'agregar canal',
            data: {
                id: canales.length,
                nombre
            }
        }
        this.ws.send(JSON.stringify(msg)); */
    }
    setCanal(canalActivo){
        this.setState({canalActivo});
        this.socket.emit('desubscribir mensaje');
        this.setState({mensajes:[]});
        // obtener canales de mensajes
        this.socket.emit('subscribir mensaje', { canalId: canalActivo.id });
    }
    setNombreUsuario(nombre){
        /*let {usuarios} = this.state;
        usuarios.push({id:usuarios.length, nombre});
        this.setState({usuarios});*/
        // enviar al servidor
        this.socket.emit('editar usuario', { nombre });
    }
    agregarMensaje(body){
        /*let {mensajes, usuarios} = this.state;
        let createdAt = new Date;
        let autor = usuarios.length > 0 ? usuarios[0].nombre: 'anonimo';
        mensajes.push({id:mensajes.length, body, createdAt, autor});
        this.setState({mensajes});*/
        // enviar al servidor
        let {canalActivo} = this.state;
        this.socket.emit('agregar mensaje', {canalId: canalActivo.id, body})
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