import RegisterForm from "../components/RegisterForm";
import { withFormik } from 'formik';

export default withFormik({
    mapPropsToValues: () => ({
        username: '',
        email: '',
        password: '',
        repeatPassword: ''
    }),
    validate: values => {
        const errors = {};
        if (!values.username) {
            errors.username = 'Введите ваше имя';
        }
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
        if (!values.password) {
            errors.repeatPassword = 'Введите ваш пароль';
        } else if (values.password !== values.repeatPassword) {
            errors.repeatPassword = 'Пароли не совпадают';
        }
        return errors;
    },

    onSubmit: (values, { setSubmitting }) => {
        setTimeout(() => {
            alert(JSON.stringify(values, null, 2));
            setSubmitting(false);
        }, 400);
    }
})(RegisterForm);