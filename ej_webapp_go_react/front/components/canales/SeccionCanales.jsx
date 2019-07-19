import React, { Component } from 'react';
import PropTypes from 'prop-types';
import FormCanal from './FormCanal.jsx';
import ListaCanales from './ListaCanales.jsx';

class SeccionCanales extends Component {
    render() {
        return (
            <div>
                <ListaCanales {...this.props}/>
                <FormCanal {...this.props}/>
            </div>
        )
    }
}

SeccionCanales.propTypes = {
    canales: PropTypes.array.isRequired,
    setCanal: PropTypes.func.isRequired,    
    agregarCanal: PropTypes.func.isRequired
}

export default SeccionCanales;