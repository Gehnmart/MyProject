import React from 'react';
import { LoginForm, RegisterForm } from 'modules';
import './Auth.scss';
import { useLocation } from 'react-router-dom';

const Auth = () => {
    const location = useLocation();
    return (
        <section className='auth'>
            <div className="auth__content">
                {
                location.pathname === '/login' ? <LoginForm /> :
                location.pathname === '/' ? <LoginForm/> :
                location.pathname === '/register' ? <RegisterForm/> :
                null 
                }    
            </div>
        </section>
    );
}

export default Auth;
