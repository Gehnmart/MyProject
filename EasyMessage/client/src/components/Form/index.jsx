import React from 'react';
import PropTypes from 'prop-types'
import classNames from 'classnames'
import { Form as BaseForm } from 'antd';
import './Button.scss';

const Form = ({ children, ...props }) => {
    return (
        <BaseForm
            {...props}
            className={classNames(
                'form',
                props.className,
            )}
        >{children}</BaseForm>
    );
}

Form.propTypes = {
    className: PropTypes.string,
};

export default Form;