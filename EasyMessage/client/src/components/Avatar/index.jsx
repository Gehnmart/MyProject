import React from 'react';
import name_to_rgb from "../../helpers/name_to_rgb";

import './Avatar.scss';
const Avatar = ({user, into_dialog}) => {
    const classname = into_dialog ? 'avatar avatar--dialog avatar--micro-letter' : 'avatar';
    if(user.avatar) {
        return (
                <img className={classname} src={user.avatar} alt={`avatar ${user.name}`}/>
        );
    }else{
        const biggest_letter = user.name.charAt(0).toUpperCase();
        const colors = name_to_rgb(user.name);
        return (
            <div className={ classname + " avatar--letter"} style={
                {background: `linear-gradient(135deg, ${colors[0]} 0%, ${colors[1]} 96.2%)`}
            }>
                {biggest_letter}
            </div>
        );
    }
};

export default Avatar;