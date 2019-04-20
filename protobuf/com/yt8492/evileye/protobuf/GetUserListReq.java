// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: protobuf/api.proto

package com.yt8492.evileye.protobuf;

/**
 * Protobuf type {@code evileye.GetUserListReq}
 */
public  final class GetUserListReq extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:evileye.GetUserListReq)
    GetUserListReqOrBuilder {
private static final long serialVersionUID = 0L;
  // Use GetUserListReq.newBuilder() to construct.
  private GetUserListReq(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GetUserListReq() {
    limit_ = 0L;
    offset_ = 0L;
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private GetUserListReq(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 8: {

            limit_ = input.readInt64();
            break;
          }
          case 16: {

            offset_ = input.readInt64();
            break;
          }
          default: {
            if (!parseUnknownFieldProto3(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.yt8492.evileye.protobuf.EeProto.internal_static_evileye_GetUserListReq_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.yt8492.evileye.protobuf.EeProto.internal_static_evileye_GetUserListReq_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.yt8492.evileye.protobuf.GetUserListReq.class, com.yt8492.evileye.protobuf.GetUserListReq.Builder.class);
  }

  public static final int LIMIT_FIELD_NUMBER = 1;
  private long limit_;
  /**
   * <code>int64 limit = 1;</code>
   */
  public long getLimit() {
    return limit_;
  }

  public static final int OFFSET_FIELD_NUMBER = 2;
  private long offset_;
  /**
   * <code>int64 offset = 2;</code>
   */
  public long getOffset() {
    return offset_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (limit_ != 0L) {
      output.writeInt64(1, limit_);
    }
    if (offset_ != 0L) {
      output.writeInt64(2, offset_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (limit_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(1, limit_);
    }
    if (offset_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(2, offset_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.yt8492.evileye.protobuf.GetUserListReq)) {
      return super.equals(obj);
    }
    com.yt8492.evileye.protobuf.GetUserListReq other = (com.yt8492.evileye.protobuf.GetUserListReq) obj;

    boolean result = true;
    result = result && (getLimit()
        == other.getLimit());
    result = result && (getOffset()
        == other.getOffset());
    result = result && unknownFields.equals(other.unknownFields);
    return result;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + LIMIT_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getLimit());
    hash = (37 * hash) + OFFSET_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getOffset());
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.yt8492.evileye.protobuf.GetUserListReq parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.yt8492.evileye.protobuf.GetUserListReq prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code evileye.GetUserListReq}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:evileye.GetUserListReq)
      com.yt8492.evileye.protobuf.GetUserListReqOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.yt8492.evileye.protobuf.EeProto.internal_static_evileye_GetUserListReq_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.yt8492.evileye.protobuf.EeProto.internal_static_evileye_GetUserListReq_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.yt8492.evileye.protobuf.GetUserListReq.class, com.yt8492.evileye.protobuf.GetUserListReq.Builder.class);
    }

    // Construct using com.yt8492.evileye.protobuf.GetUserListReq.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      limit_ = 0L;

      offset_ = 0L;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.yt8492.evileye.protobuf.EeProto.internal_static_evileye_GetUserListReq_descriptor;
    }

    @java.lang.Override
    public com.yt8492.evileye.protobuf.GetUserListReq getDefaultInstanceForType() {
      return com.yt8492.evileye.protobuf.GetUserListReq.getDefaultInstance();
    }

    @java.lang.Override
    public com.yt8492.evileye.protobuf.GetUserListReq build() {
      com.yt8492.evileye.protobuf.GetUserListReq result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.yt8492.evileye.protobuf.GetUserListReq buildPartial() {
      com.yt8492.evileye.protobuf.GetUserListReq result = new com.yt8492.evileye.protobuf.GetUserListReq(this);
      result.limit_ = limit_;
      result.offset_ = offset_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return (Builder) super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return (Builder) super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return (Builder) super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return (Builder) super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return (Builder) super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return (Builder) super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.yt8492.evileye.protobuf.GetUserListReq) {
        return mergeFrom((com.yt8492.evileye.protobuf.GetUserListReq)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.yt8492.evileye.protobuf.GetUserListReq other) {
      if (other == com.yt8492.evileye.protobuf.GetUserListReq.getDefaultInstance()) return this;
      if (other.getLimit() != 0L) {
        setLimit(other.getLimit());
      }
      if (other.getOffset() != 0L) {
        setOffset(other.getOffset());
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      com.yt8492.evileye.protobuf.GetUserListReq parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (com.yt8492.evileye.protobuf.GetUserListReq) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private long limit_ ;
    /**
     * <code>int64 limit = 1;</code>
     */
    public long getLimit() {
      return limit_;
    }
    /**
     * <code>int64 limit = 1;</code>
     */
    public Builder setLimit(long value) {
      
      limit_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 limit = 1;</code>
     */
    public Builder clearLimit() {
      
      limit_ = 0L;
      onChanged();
      return this;
    }

    private long offset_ ;
    /**
     * <code>int64 offset = 2;</code>
     */
    public long getOffset() {
      return offset_;
    }
    /**
     * <code>int64 offset = 2;</code>
     */
    public Builder setOffset(long value) {
      
      offset_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 offset = 2;</code>
     */
    public Builder clearOffset() {
      
      offset_ = 0L;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFieldsProto3(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:evileye.GetUserListReq)
  }

  // @@protoc_insertion_point(class_scope:evileye.GetUserListReq)
  private static final com.yt8492.evileye.protobuf.GetUserListReq DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.yt8492.evileye.protobuf.GetUserListReq();
  }

  public static com.yt8492.evileye.protobuf.GetUserListReq getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GetUserListReq>
      PARSER = new com.google.protobuf.AbstractParser<GetUserListReq>() {
    @java.lang.Override
    public GetUserListReq parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new GetUserListReq(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<GetUserListReq> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GetUserListReq> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.yt8492.evileye.protobuf.GetUserListReq getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}
