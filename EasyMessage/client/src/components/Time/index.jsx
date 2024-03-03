import PropTypes from 'prop-types'
import { formatDistanceToNow } from "date-fns";
import { ru } from 'date-fns/locale'
import './Time.scss';

const Time = ({date}) => 
    formatDistanceToNow(new Date(date), { addSuffix: true, locale: ru });

Time.propTypes = {
    date: PropTypes.string
};

export default Time;
