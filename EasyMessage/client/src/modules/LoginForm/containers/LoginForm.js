import { withFormik } from 'formik';
import LoginForm from '../components/LoginForm';

export default withFormik({
    mapPropsToValues: () => ({
        email: '',
        password: '',
    }),
    validate: values => {
        const errors = {};
        if (!values.email || values.email.length < 6) {
            errors.email = 'Введите ваш email';
        } else if (
            !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i.test(values.email)
        ) {
            errors.email = 'Некорректный email адрес';
        }
        if (!values.password) {
            errors.password = 'Введите ваш пароль';
        } else if (values.password.length < 8) {
            errors.password = 'Слишком короткий пароль';
        } else if (values.password.length > 32) {
            errors.password = 'Слишком длинный пароль';
        } else if (!/^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/.test(values.password)) {
            errors.password = 'Пароль должен содержать буквы и цифры';
        }
        return errors;
    },

    onSubmit: (values, { setSubmitting }) => {
        setTimeout(() => {
            alert(JSON.stringify(values, null, 2));
            setSubmitting(false);
        }, 400);
    }
})(LoginForm);