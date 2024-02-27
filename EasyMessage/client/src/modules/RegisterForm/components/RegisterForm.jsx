import React from 'react';
import { Form, Input } from 'antd';
import { Button, Block } from 'components';
import { SmileOutlined, LockOutlined, MailOutlined } from '@ant-design/icons';
import { Link } from 'react-router-dom'
import { Formik } from 'formik';

const RegisterForm = props => {
    const {
        handleSubmit,
        handleChange,
        handleBlur,
        values,
        errors,
        touched,
        isSubmitting
    } = props;
    return (
        <div>
            <div className='auth__top'>
                <h2>Создайте аккаунт</h2>
                <p>Пожалуйста, создайте новый аккаунт</p>
            </div>
            <Block>
                <Formik>
                    <Form onSubmit={handleSubmit} className="login-form">
                        <Form.Item
                            validateStatus={errors.username && touched.username ? "error" : "success"}
                            help={errors.username && touched.username && errors.username}
                            hasFeedback
                        >
                            <Input
                                size='large'
                                prefix={<SmileOutlined />}
                                id='username'
                                type='username'
                                name="username"
                                onChange={handleChange}
                                onBlur={handleBlur}
                                value={values.username}
                                placeholder="Username"
                            />
                        </Form.Item>
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
                                placeholder="Email"
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
                        <Form.Item
                            validateStatus={errors.repeatPassword && touched.repeatPassword ? "error" : "success"}
                            help={errors.repeatPassword && touched.repeatPassword && errors.repeatPassword}
                            hasFeedback
                        >
                            <Input
                                size='large'
                                prefix={<LockOutlined />}
                                id='repeatPassword'
                                type="password"
                                name="repeatPassword"
                                onChange={handleChange}
                                onBlur={handleBlur}
                                value={values.repeatPassword}
                                placeholder="Repeat password"
                            />
                        </Form.Item>
                        <Form.Item>
                            <Button type="primary"
                                htmlType="submit"
                                size="large"
                                className="button"
                                disabled={true && (isSubmitting || Object.keys(errors).length > 0)}
                            >
                                Создать аккаунт
                            </Button>
                        </Form.Item>
                        <Form.Item>
                            <p></p>
                            <Link className='auth__login-link' to='/'>Уже есть аккаунт?
                                Войти</Link>
                        </Form.Item>
                    </Form>
                </Formik>
            </Block>
        </div>
    );
}

export default RegisterForm;
