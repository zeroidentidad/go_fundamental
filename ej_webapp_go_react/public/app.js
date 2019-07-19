let canales = [
    {nombre: 'Soporte Tecnico'},
    {nombre: 'Programación Web'}
]

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

ReactDOM.render(<ListaCanales canales={canales}/>, document.getElementById('app'));