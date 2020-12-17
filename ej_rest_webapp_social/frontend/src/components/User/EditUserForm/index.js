import React, {useState} from "react";
import { Form, Button, Row, Col, Spinner } from "react-bootstrap";

import "./EditUserForm.scss";

export default function EditUserForm(props) {
const [loading, setLoading] = useState(false);

const onSubmit = async (e) => {
    e.preventDefault();
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
                    />
                    </Col>
                    <Col>
                    <Form.Control
                        type="text"
                        placeholder="Apellidos"
                        name="apellidos"
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
                />
                </Form.Group>

                <Form.Group>
                <Form.Control
                    type="text"
                    placeholder="Sitio web"
                    name="sitioWeb"
                />
                </Form.Group>

                <Button className="btn-submit" variant="primary" type="submit">
                {loading && <Spinner animation="border" size="sm" />} Actualizar
                </Button>
            </Form>
        </div>
    )
}