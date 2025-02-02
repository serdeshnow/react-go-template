import ReactLogo from '@/shared/assets/svg/react.svg?react';
import { useState } from 'react';
import viteLogo from '/vite.svg';
import s from './Page.module.scss';
import { env } from '@/shared/lib/env.ts';

export const HomePage = () => {
  const [count, setCount] = useState<number>(0);

  return (
    <>
      <div className={`flex alignCenter justifyCenter`}>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className={s.logo} alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <ReactLogo className={`${s.logo} ${s.react}`} />
        </a>
      </div>
      <div className={`flexColumn`}>
        <p>APP_NAME: {env.APP_NAME}</p>
        <p>API_URL: {env.API_URL}</p>
        <p>ENV: {env.ENV}</p>
      </div>
      <h1>Vite + React</h1>
      <div className={s.card}>
        <button onClick={() => setCount((count) => count + 1)}>count is {count}</button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className={s.readTheDocs}>Click on the Vite and React logos to learn more</p>
    </>
  );
};
