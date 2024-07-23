import React, { FC, Suspense, useMemo } from 'react';
import { Redirect, Route, Switch } from 'react-router-dom';

import { Loader, Page404 } from 'ui';
import { ProtectedRoute } from './ProtectedRoute';

import { useProfile } from 'hooks';

import { getInitialRoute } from 'helpers';
import { ROUTES } from 'consts';

import { Routes } from './Routes';

export const SwitchRoutes: FC = () => {
  const {
    profile: { role, accountingType },
  } = useProfile();

  const initialRoute = useMemo(() => getInitialRoute(role, accountingType), [role, accountingType]);

  return (
    <Suspense fallback={<Loader type='start' />}>
      <Switch>
        {Routes.map((route, i) => {
          const { path, isProtected, component } = route;

          if (isProtected) {
            return <ProtectedRoute key={i} {...route} component={component} />;
          } else {
            <Route path={path} component={component} />;
          }
        })}

        <Route exact path={ROUTES.SYSTEM.ROOT}>
          <Redirect to={initialRoute} />
        </Route>

        <Route component={Page404} />
      </Switch>
      {/* <Redirect from={ROUTES.SYSTEM.ROOT} to={initialRoute} /> */}
    </Suspense>
  );
};
