import React from 'react';
import { Container } from 'semantic-ui-react';
import Header from './Header';
import Footer from './Footer';

export default props => {
    return (
        <Container>
            <div className="App">
                <Header />
                    { props.children }
                <Footer />
            </div>
        </Container>
    );
};