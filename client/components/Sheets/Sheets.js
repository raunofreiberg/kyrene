// @flow
import React, { Fragment } from 'react';
import moment from 'moment';
import { Link } from 'react-router-dom';
import { graphql } from 'react-apollo';
import classNames from 'classnames';

import * as styles from './Sheets.scss';
import { sheetsQuery } from '../../graphql/queries/Sheets.graphql';

import Avatar from '../Avatar';
import Spinner from '../Spinner';

type Props = {
  data: {
      sheets: Array<Object>,
      loading: boolean,
  }
};

const Sheets = ({ data: { sheets, loading } }: Props) => {
    if (loading) {
        return <Spinner />;
    }

    return (
        <div className={classNames('container', styles.sheetsWrapper)}>
            {sheets && sheets.map((sheet, idx) => {                
                const { id, name, createdAt, user: { username } } = sheet;
                
                return (
                    <Fragment key={id}>
                        <Link to={`/sheet/${id}`} className={classes} title={name}>
                            <div className={styles.sheetDetails}>
                                <h1>{name}</h1>
                                <p className="note">Created {createdAtTimestamp}</p>
                            </div>
                            <Avatar username={username} />
                        </Link>
                    </Fragment>
                )
            })}
        </div>
    );
};

export default graphql(sheetsQuery)(Sheets);
