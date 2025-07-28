package auth

type UseCase interface {
	Register(payload *RegisterRequestDTO) (*RegisterResponseDTO, error)
}
