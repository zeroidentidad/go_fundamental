import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Usuario from './Usuario.jsx';

class ListaUsuarios extends Component {
    render() {
        return (
            <ul>
            {
                this.props.usuarios.map(usuario=>{
                    return(
                        <Usuario key={usuario.id} usuario={usuario}/>
                    )
                })
            }
            </ul>
        )
    }
}

ListaUsuarios.propTypes = {
    usuarios: PropTypes.array.isRequired
}

export default ListaUsuarios;