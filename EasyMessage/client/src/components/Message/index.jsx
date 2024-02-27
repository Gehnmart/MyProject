import React from 'react';
import PropTypes from 'prop-types'
import { formatDistanceToNow } from "date-fns";
import { ru } from 'date-fns/locale'
import './Message.scss';

const Message = ({avatar, text, date, user, is_me}) => {
    const messageClass = is_me ? 'message message--me' : 'message';

    return (
        <div className={messageClass}>
            <div className="message__avatar">
                <img src={avatar} alt={`avatar ${user.name}`} />
            </div>
            <div className="message__content">
                <div className="message__bubble">
                    <p className='message__text'>{text}</p>
                </div>
                <span className='message__date'>
                    {formatDistanceToNow(new Date(date), { addSuffix: true, locale: ru })}
                </span>
            </div>
        </div>
    );
}

Message.defaultProps = {
    user: {}
};

Message.propTypes = {
    avatar: PropTypes.string,
    text: PropTypes.string,
    date: PropTypes.string,
    user: PropTypes.object
};

export default Message;
