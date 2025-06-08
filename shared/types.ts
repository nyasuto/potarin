export interface Suggestion {
  id: string;
  title: string;
  description: string;
}

export interface Position {
  lat: number;
  lng: number;
}

export interface Route {
  title: string;
  description: string;
  position: Position;
}

export interface Detail {
  summary: string;
  routes: Route[];
}
