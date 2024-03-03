import React from 'react';

import './Dialogitem.scss';
import { Avatar } from 'components';

const DialogItem = ({user, message}) => {
    return (
    <div className='dialogs__item'>
        <div className="dialogs__item-avatar">
            <Avatar user={user} />
        </div>
        <div className="dialogs__item-info">
            <div className="dialogs__item-info-top">
                <p>{user.name}</p>
            </div>
            <div className="dialogs__item-info-bottom">
                {message}
            </div>
        </div>
    </div>
    )
}


export default DialogItem;
