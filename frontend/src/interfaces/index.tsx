export interface SignUpRequest {
  username: string,
  password: string
}

export interface SignInRequest {
  username: string,
  password: string
}

export interface ListNoteRequest {
  page: number,
  size: number
}

export interface NewNote {
  title: string,
  text: string
}

export interface NoteOverview {
  id: number,
  title: string,
  text: string,
  createdAt: Date,
  updatedAt: Date
}

export interface ListNoteResponse {
  notes: NoteOverview[],
  isLast: boolean,
  total: number
}
