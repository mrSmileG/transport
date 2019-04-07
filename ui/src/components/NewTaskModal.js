import React, { Component } from 'react';
import { Button, Modal, Header, Input, Message, Form } from 'semantic-ui-react';

class NewTaskModal extends Component {
    state = {
        providersCount: '',
        consumersCount: '',
        errorMessage: '',
        loading: false,
        isOpen: false
    };

    onSubmit = async event => {
        event.preventDefault();
        const {
            providersCount,
            consumersCount
        } = this.state;

        this.setState({
            errorMessage: '',
            loading: true
        });

        try {
            this.props.updateData(providersCount, consumersCount);
            this.close();
        } catch(error) {
            this.setState({
                errorMessage: error.message
            });
        }

        this.setState({
            loading: false
        });
    }

    close = () => {
        this.setState({
            isOpen: false
        });
    }

    open = () => {
        this.setState({
            isOpen: true
        });
    }

    render() {
        const {
            providersCount,
            consumersCount,
            errorMessage,
            loading,
            isOpen
        } = this.state;

        return (
            <div>
            <Modal trigger={<Button 
                    primary
                    floated="right"
                    content="New Task"
                    icon="add circle"
                    style={{display: 'block', marginTop: '10px'}}
                    onClick={this.open}
                />}
                size="mini"
                closeIcon
                onClose={this.close}
                open={isOpen}
                >
                <Header 
                    icon="edit" 
                    content="New Task" 
                    />
                <Modal.Content>

                    <Form onSubmit={this.onSubmit} error={!!errorMessage}>
                        <Form.Field>
                            <label>Providers Count</label>
                            <Input
                                value={providersCount}
                                onChange={event =>
                                    this.setState({
                                        providersCount: event.target.value
                                    })
                                }
                            />
                        </Form.Field>  
                        <Form.Field>
                            <label>Consumers Count</label>
                            <Input
                                value={consumersCount}
                                onChange={event =>
                                    this.setState({
                                        consumersCount: event.target.value
                                    })
                                }
                            />
                        </Form.Field>                          
                        <Message error header="Oops!" content={errorMessage} />
                        
                        <div style={{float: 'right', margin: '20px 0 30px 0'}}>
                            <Button 
                                negative 
                                onClick={this.close} 
                                icon="cancel"
                                content="Cancel"
                                style={{marginRight: '10px'}}
                            />
                            <Button 
                                positive 
                                icon="check" 
                                content="Apply"
                                loading={loading} 
                                disabled={loading} 
                            />
                        </div>
                    </Form>
                </Modal.Content>
            </Modal>
            

            </div>
        );
    }
}

export default NewTaskModal;