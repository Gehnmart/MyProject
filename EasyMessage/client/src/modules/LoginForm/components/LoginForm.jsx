import React from 'react';
import { Form, Input } from 'antd';
import { Button, Block } from 'components';
import { MailOutlined, LockOutlined } from '@ant-design/icons';
import { Link } from 'react-router-dom';
import { Formik } from 'formik';

const LoginForm = props => {
    const {
        values,
        touched,
        errors,
        handleChange,
        handleBlur,
        handleSubmit,
        isSubmitting
    } = props;
    return (
        <div>
            <div className='auth__top'>
                <h2>Войти в аккаунт</h2>
                <p>Пожалуйста, войдите в свой аккаунт</p>
            </div>
            <Block>
                <Formik>
                    <Form onSubmit={handleSubmit} className="login-form">
                        <Form.Item
                            validateStatus={errors.email && touched.email ? "error" : "success"}
                            help={errors.email && touched.email && errors.email}
                            hasFeedback
                        >
                            <Input
                                size='large'
                                prefix={<MailOutlined />}
                                id='email'
                                type="email"
                                name="email"
                                onChange={handleChange}
                                onBlur={handleBlur}
                                value={values.email}
                                placeholder="E-Mail"
                            />
                        </Form.Item>
                        <Form.Item
                            validateStatus={errors.password && touched.password ? "error" : "success"}
                            help={errors.password && touched.password && errors.password}
                            hasFeedback
                        >
                            <Input
                                size='large'
                                prefix={<LockOutlined />}
                                type="password"
                                name="password"
                                onChange={handleChange}
                                onBlur={handleBlur}
                                value={values.password}
                                placeholder="Password"
                            />
                        </Form.Item>
                        <Form.Item>
                            <Button
                                type="primary"
                                htmlType="submit"
                                size="large"
                                className="button"
                                disabled={isSubmitting || Object.keys(errors).length > 0}
                            >
                                Войти в аккаунт
                            </Button>
                        </Form.Item>
                        <Form.Item>
                            <p></p>
                            <Link className='auth__register-link' to='/register'>Зарегистрироваться</Link>
                        </Form.Item>
                    </Form>
                </Formik>
            </Block>
        </div>
    );
}

export default LoginForm;

