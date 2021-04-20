// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *ApplicationIdentifiers) SetFields(src *ApplicationIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_id":
			if len(subs) > 0 {
				return fmt.Errorf("'application_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ApplicationID = src.ApplicationID
			} else {
				var zero string
				dst.ApplicationID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ClientIdentifiers) SetFields(src *ClientIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "client_id":
			if len(subs) > 0 {
				return fmt.Errorf("'client_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ClientID = src.ClientID
			} else {
				var zero string
				dst.ClientID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceIdentifiers) SetFields(src *EndDeviceIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "device_id":
			if len(subs) > 0 {
				return fmt.Errorf("'device_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DeviceID = src.DeviceID
			} else {
				var zero string
				dst.DeviceID = zero
			}
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				newDst = &dst.ApplicationIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "dev_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevEUI = src.DevEUI
			} else {
				dst.DevEUI = nil
			}
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEUI = src.JoinEUI
			} else {
				dst.JoinEUI = nil
			}
		case "dev_addr":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_addr' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevAddr = src.DevAddr
			} else {
				dst.DevAddr = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GatewayIdentifiers) SetFields(src *GatewayIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_id":
			if len(subs) > 0 {
				return fmt.Errorf("'gateway_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.GatewayID = src.GatewayID
			} else {
				var zero string
				dst.GatewayID = zero
			}
		case "eui":
			if len(subs) > 0 {
				return fmt.Errorf("'eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.EUI = src.EUI
			} else {
				dst.EUI = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *OrganizationIdentifiers) SetFields(src *OrganizationIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization_id":
			if len(subs) > 0 {
				return fmt.Errorf("'organization_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.OrganizationID = src.OrganizationID
			} else {
				var zero string
				dst.OrganizationID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UserIdentifiers) SetFields(src *UserIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_id":
			if len(subs) > 0 {
				return fmt.Errorf("'user_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UserID = src.UserID
			} else {
				var zero string
				dst.UserID = zero
			}
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *OrganizationOrUserIdentifiers) SetFields(src *OrganizationOrUserIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {

		case "ids":
			if len(subs) == 0 && src == nil {
				dst.Ids = nil
				continue
			} else if len(subs) == 0 {
				dst.Ids = src.Ids
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "organization_ids":
					_, srcOk := src.Ids.(*OrganizationOrUserIdentifiers_OrganizationIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*OrganizationOrUserIdentifiers_OrganizationIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *OrganizationIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*OrganizationOrUserIdentifiers_OrganizationIDs).OrganizationIDs
						}
						if dstOk {
							newDst = dst.Ids.(*OrganizationOrUserIdentifiers_OrganizationIDs).OrganizationIDs
						} else {
							newDst = &OrganizationIdentifiers{}
							dst.Ids = &OrganizationOrUserIdentifiers_OrganizationIDs{OrganizationIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "user_ids":
					_, srcOk := src.Ids.(*OrganizationOrUserIdentifiers_UserIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*OrganizationOrUserIdentifiers_UserIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *UserIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*OrganizationOrUserIdentifiers_UserIDs).UserIDs
						}
						if dstOk {
							newDst = dst.Ids.(*OrganizationOrUserIdentifiers_UserIDs).UserIDs
						} else {
							newDst = &UserIdentifiers{}
							dst.Ids = &OrganizationOrUserIdentifiers_UserIDs{UserIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EntityIdentifiers) SetFields(src *EntityIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {

		case "ids":
			if len(subs) == 0 && src == nil {
				dst.Ids = nil
				continue
			} else if len(subs) == 0 {
				dst.Ids = src.Ids
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "application_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_ApplicationIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'application_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_ApplicationIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'application_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ApplicationIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_ApplicationIDs).ApplicationIDs
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_ApplicationIDs).ApplicationIDs
						} else {
							newDst = &ApplicationIdentifiers{}
							dst.Ids = &EntityIdentifiers_ApplicationIDs{ApplicationIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "client_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_ClientIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'client_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_ClientIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'client_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ClientIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_ClientIDs).ClientIDs
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_ClientIDs).ClientIDs
						} else {
							newDst = &ClientIdentifiers{}
							dst.Ids = &EntityIdentifiers_ClientIDs{ClientIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "device_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_DeviceIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'device_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_DeviceIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'device_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *EndDeviceIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_DeviceIDs).DeviceIDs
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_DeviceIDs).DeviceIDs
						} else {
							newDst = &EndDeviceIdentifiers{}
							dst.Ids = &EntityIdentifiers_DeviceIDs{DeviceIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "gateway_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_GatewayIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'gateway_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_GatewayIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'gateway_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *GatewayIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_GatewayIDs).GatewayIDs
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_GatewayIDs).GatewayIDs
						} else {
							newDst = &GatewayIdentifiers{}
							dst.Ids = &EntityIdentifiers_GatewayIDs{GatewayIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "organization_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_OrganizationIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_OrganizationIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *OrganizationIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_OrganizationIDs).OrganizationIDs
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_OrganizationIDs).OrganizationIDs
						} else {
							newDst = &OrganizationIdentifiers{}
							dst.Ids = &EntityIdentifiers_OrganizationIDs{OrganizationIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "user_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_UserIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_UserIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *UserIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_UserIDs).UserIDs
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_UserIDs).UserIDs
						} else {
							newDst = &UserIdentifiers{}
							dst.Ids = &EntityIdentifiers_UserIDs{UserIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
