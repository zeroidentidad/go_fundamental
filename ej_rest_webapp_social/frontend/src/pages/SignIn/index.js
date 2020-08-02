import React from 'react';
import { Container, Row, Col, Button } from "react-bootstrap";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch, faUsers, faComment } from "@fortawesome/free-solid-svg-icons";
import LogoTwittab from "../../assets/png/logo.png";
import LogoWhiteTwittab from "../../assets/png/logo-white.png";
import "./index.scss";

export default function SignIn () {
    return (
        <>
            <Container className="signin-signup" fluid>
                <Row>
                    <LeftComponent />
                    <RightComponent />
                </Row>
            </Container>
        </>
    )
}

function LeftComponent() {
    return (
        <Col className="signin-signup__left" xs={6}>
            <img src={LogoTwittab} alt="Twittab" />
            <div>
                <h2>
                    <FontAwesomeIcon icon={faSearch} />
                    Sigue lo que te interesa.
                </h2>
                <h2>
                    <FontAwesomeIcon icon={faUsers} />
                    Entérate de qué está hablando la gente.
                </h2>
                <h2>
                    <FontAwesomeIcon icon={faComment} />
                    Únete a la conversación.
                </h2>
            </div>
        </Col>
    );
}

function RightComponent(props) {

    return (
        <Col className="signin-signup__right" xs={6}>
            <div>
                <img src={LogoWhiteTwittab} alt="Twittab" />
                <h2>Mira lo que está alborotando a la chocada en este momento</h2>
                <h3>Únete a Twittab ahora</h3>
                <Button
                    variant="primary"
                    onClick={() => alert('kepedo')}
                >
                    Regístrate
                </Button>
                <Button
                    variant="outline-primary"
                    onClick={() => alert('kepedo')}
                >
                    Iniciar sesión
                </Button>
            </div>
        </Col>
    );
}