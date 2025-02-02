#!/bin/bash
rm -rf frontend/web/public
cd frontend/react
npm run build
cd ..
mv react/out web/public