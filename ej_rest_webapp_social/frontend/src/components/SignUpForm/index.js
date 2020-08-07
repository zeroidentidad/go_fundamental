import React, { useState } from "react";
import { Row, Col, Form, Button, Spinner } from "react-bootstrap";

import "./SignUpForm.scss";

export default function SignUpForm(props) {
  const { setShowModal } = props;

  const onSubmit = e => {
    e.preventDefault();
    
    setShowModal(false);
  };

  return (
    <div className="sign-up-form">
      <h2>Crear cuenta</h2>
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
            type="email"
            placeholder="Correo electronico"
            name="email"
          />
        </Form.Group>
        <Form.Group>
          <Row>
            <Col>
              <Form.Control
                type="password"
                placeholder="Contraseña"
                name="password"
              />
            </Col>
            <Col>
              <Form.Control
                type="password"
                placeholder="Confirmar contraseña"
                name="repeatPassword"
              />
            </Col>
          </Row>
        </Form.Group>

        <Button variant="primary" type="submit">
          Registrarme
        </Button>
      </Form>
    </div>
  );
}