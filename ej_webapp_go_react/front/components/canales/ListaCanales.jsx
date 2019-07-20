import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Canal from './Canal.jsx';

class ListaCanales extends Component {
    render() {
        return (
            <ul>
                {this.props.canales.map(can=>{
                    return <Canal canal={can}
                    key={can.id}
                    /*setCanal={this.props.setCanal}
                    canalActivo={this.props.canalActivo}*/
                    {...this.props}
                    />
                })}
            </ul>
        )
    }
}

ListaCanales.propTypes ={
    canales: PropTypes.array.isRequired,
    setCanal: PropTypes.func.isRequired,
    canalActivo: PropTypes.object.isRequired
}

export default ListaCanales;