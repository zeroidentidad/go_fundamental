
class Canal extends React.Component{
    onClick(){
        console.log('Click en', this.props.nombre)
    }
    render(){
        return(
            <li onClick={this.onClick.bind(this)}>{this.props.nombre}</li>
        )
    }
}

class ListaCanales extends React.Component{
    render(){
        return(
            <ul>
                {this.props.canales.map(canal=>{
                    return (
                        <Canal nombre={canal.nombre} />
                    )
                })
                }
            </ul>
        )
    }
}

class FormCanal extends React.Component{
    constructor(props){
        super(props);
        this.state={};
    }
    onChange(e) {
        this.setState({
            nombreCanal: e.target.value
        });
        //console.log(e.target.value);
    }
    onSubmit(e) {
        let { nombreCanal } = this.state;
        console.log(nombreCanal);
        this.setState({
            nombreCanal: ''
        });
        this.props.agregarCanal(nombreCanal)
        e.preventDefault();
    }    
    render(){
        return(
            <form onSubmit={this.onSubmit.bind(this)}>
                <input type='text' onChange={this.onChange.bind(this)} value={this.state.nombreCanal}></input>
            </form>
        )
    }
}

class SeccionCanal extends React.Component{
    constructor(props){
        super(props);
        this.state={
            canales:[
                { nombre: 'Soporte Tecnico' },
                { nombre: 'Programación Web' }
            ]
        };
    }
    agregarCanal(nombre){
        let {canales} = this.state;
        canales.push({nombre: nombre});
        this.setState({
            canales: canales
        });
    }
    render(){
        return(
            <div>
                <ListaCanales canales={this.state.canales} />
                <FormCanal agregarCanal={this.agregarCanal.bind(this)}/>
            </div>
        )
    }
}

ReactDOM.render(<SeccionCanal/>, document.getElementById('app'));