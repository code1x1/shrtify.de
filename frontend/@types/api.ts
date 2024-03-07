import { z } from "zod";

export const MetaSchema = z.object({
  status: z.number(),
  error: z.string().nullable(),
  more: z.any(),
})
export type Meta = z.infer<typeof MetaSchema>

export const CreateShortcodeRequestSchema = z.object({
  url: z.string(),
  code: z.string(),
})
export type CreateShortcodeRequest = z.infer<typeof CreateShortcodeRequestSchema>

export const CreateShortcodeRequestPayloadSchema = z.object({
  request: CreateShortcodeRequestSchema,
})
export type CreateShortcodeRequestPayload = z.infer<typeof CreateShortcodeRequestPayloadSchema>

export const CreateShortcodeResponseSchema = z.object({
  shortUrl: z.string(),
})
export type CreateShortcodeResponse = z.infer<typeof CreateShortcodeResponseSchema>

export const CreateShortcodeResponsePayloadSchema = z.object({
  meta: MetaSchema,
  response: CreateShortcodeResponseSchema,
})
export type CreateShortcodeResponsePayload = z.infer<typeof CreateShortcodeResponsePayloadSchema>

export const GetShortcodesResponseSchema = z.object({
  shortUrl: z.string(),
  originalUrl: z.string(),
})
export type GetShortcodesResponse = z.infer<typeof GetShortcodesResponseSchema>

export const GetShortcodesResponsePayloadSchema = z.object({
  meta: MetaSchema,
  response: GetShortcodesResponseSchema.array().nullable(),
})
export type GetShortcodesResponsePayload = z.infer<typeof GetShortcodesResponsePayloadSchema>

export const GetShortcodesMetaSchema = z.object({
  pages: z.number(),
})
export type GetShortcodesMeta = z.infer<typeof GetShortcodesMetaSchema>

export const ApiSchema = z.object({
  Meta: MetaSchema,
  CreateShortcodeRequestPayload: CreateShortcodeRequestPayloadSchema,
  CreateShortcodeResponsePayload: CreateShortcodeResponsePayloadSchema,
  GetShortcodesResponsePayload: GetShortcodesResponsePayloadSchema,
  GetShortcodesMeta: GetShortcodesMetaSchema,
})
export type Api = z.infer<typeof ApiSchema>

