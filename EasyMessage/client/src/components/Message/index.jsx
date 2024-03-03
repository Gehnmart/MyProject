import React from 'react';
import PropTypes from 'prop-types'
import Time from 'components/Time'
import './Message.scss';
import Avatar from 'components/Avatar';

const Message = ({avatar, text, date, user, is_me}) => {
    const messageClass = is_me ? 'message message--me' : 'message';

    return (
        <div className={messageClass}>
            <div className="message__avatar">
                <Avatar user={user} into_dialog={true} />
            </div>
            <div className="message__content">
                <div className="message__bubble">
                    <p className='message__text'>{text}</p>
                </div>
                <span className='message__date'>
                    <Time date={date}></Time>
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
