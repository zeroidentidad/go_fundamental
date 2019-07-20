import React, { Component } from 'react';
import PropTypes from 'prop-types';

class Usuario extends Component {
    render() {
        return (
            <li>
                {this.props.usuario.nombre}
            </li>
        )
    }
}

Usuario.propTypes = {
    usuario: PropTypes.object.isRequired
}

export default Usuario;
