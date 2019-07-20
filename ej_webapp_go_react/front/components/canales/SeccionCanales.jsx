import React, { Component } from 'react';
import PropTypes from 'prop-types';
import FormCanal from './FormCanal.jsx';
import ListaCanales from './ListaCanales.jsx';

class SeccionCanales extends Component {
    render() {
        return (
            <div className='support panel panel-primary'>
                <div className='panel-heading'>
                    <strong>Canales</strong>
                </div>
                <div className='panel-body channels'>
                    <ListaCanales {...this.props} />
                    <FormCanal {...this.props} />
                </div>                
            </div>
        )
    }
}

SeccionCanales.propTypes = {
    canales: PropTypes.array.isRequired,
    setCanal: PropTypes.func.isRequired,    
    agregarCanal: PropTypes.func.isRequired,
    canalActivo: PropTypes.object.isRequired
}

export default SeccionCanales;