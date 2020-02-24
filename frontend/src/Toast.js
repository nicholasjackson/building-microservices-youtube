import React from 'react';
import Toast from 'react-bootstrap/Toast';

class Toaster extends React.Component {

    constructor(props) {
        super(props);
        this.state = {show: false};
        this.hide = this.hide.bind(this);
    }

    hide(){
        this.setState({show: false});
    }

    componentDidUpdate(prevProps) {
        if (this.props !== prevProps) {
            this.setState({show: this.props.show});
        }
    }

    render() {
        return(
              <Toast onClose={this.hide} show={this.state.show} delay={3000} autohide>
                <Toast.Header>
                  <strong className="mr-auto">File Upload</strong>
                </Toast.Header>
                <Toast.Body>{this.props.message}</Toast.Body>
              </Toast>
        )
    }
}

export default Toaster;