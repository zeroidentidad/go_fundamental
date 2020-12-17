import React, {useState} from "react";
import { Form, Button, Row, Col, Spinner } from "react-bootstrap";
import DatePicker from "react-datepicker";
import es from "date-fns/locale/es";

import "./EditUserForm.scss";

export default function EditUserForm(props) {
    const { user, setShowModal } = props;
    const [loading, setLoading] = useState(false);
    const [formData, setFormData] = useState(initialFormValue(user));
    
    const onSubmit = async (e) => {
        e.preventDefault();
        console.log(formData);
    }

    return (
        <div className="edit-user-form">
            <Form onSubmit={onSubmit}>
                <Form.Group>
                <Row>
                    <Col>
                    <Form.Control
                        type="text"
                        placeholder="Nombre"
                        name="nombre"
                        defaultValue={formData.nombre}
                    />
                    </Col>
                    <Col>
                    <Form.Control
                        type="text"
                        placeholder="Apellidos"
                        name="apellidos"
                        defaultValue={formData.apellidos}
                    />
                    </Col>
                </Row>
                </Form.Group>

                <Form.Group>
                <Form.Control
                    as="textarea"
                    row="3"
                    placeholder="Agrega tu biografía"
                    type="text"
                    name="biografia"
                    defaultValue={formData.biografia}
                />
                </Form.Group>

                <Form.Group>
                <Form.Control
                    type="text"
                    placeholder="Sitio web"
                    name="sitioWeb"
                    defaultValue={formData.sitioweb}
                />
                </Form.Group>

                <Form.Group>
                <DatePicker
                    placeholder="Fecha de nacimiento"
                    locale={es}
                    selected={new Date(formData.fechaNacimiento)}
                />
                </Form.Group>

                <Button className="btn-submit" variant="primary" type="submit">
                {loading && <Spinner animation="border" size="sm" />} Actualizar
                </Button>
            </Form>
        </div>
    )
}

const initialFormValue = (user) => ({
    nombre: user.nombre || "",
    apellidos: user.apellidos || "",
    biografia: user.biografia || "",
    ubicacion: user.ubicacion || "",
    sitioweb: user.sitioweb || "",
    fechaNacimiento: user.fechaNacimiento || "",
});