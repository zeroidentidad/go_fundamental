import React, { Component } from 'react';
import PropTypes from 'prop-types';
import ListaUsuarios from './ListaUsuarios.jsx';
import FormUsuario from './FormUsuario.jsx';

class SeccionUsuarios extends Component {
    render() {
        return (
            <div className='support panel panel-primary'>
                <div className='panel-heading'>
                    <strong>Usuarios</strong>
                </div>
                <div className='panel-body users'>
                    <ListaUsuarios {...this.props}/>
                    <FormUsuario {...this.props} />
                </div>
            </div>
        )
    }
}

SeccionUsuarios.propTypes = {
    usuarios: PropTypes.array.isRequired,
    setNombreUsuario: PropTypes.func.isRequired
}


export default SeccionUsuarios;