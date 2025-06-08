'use client'

import dynamic from 'next/dynamic'
import { Route } from 'potarin-shared/types'

const MapClient = dynamic(() => import('./MapClient'), { ssr: false })

export default function Map({ routes }: { routes: Route[] }) {
  return <MapClient routes={routes} />
}
