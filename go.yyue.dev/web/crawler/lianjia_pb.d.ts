import * as jspb from "google-protobuf"

export class HouseSummary extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getHouseNo(): string;
  setHouseNo(value: string): void;

  getUrl(): string;
  setUrl(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  getTotalPrice(): string;
  setTotalPrice(value: string): void;

  getUnitPrice(): string;
  setUnitPrice(value: string): void;

  getPlot(): string;
  setPlot(value: string): void;

  getRegion(): string;
  setRegion(value: string): void;

  getLayout(): string;
  setLayout(value: string): void;

  getArea(): string;
  setArea(value: string): void;

  getFace(): string;
  setFace(value: string): void;

  getDecoration(): string;
  setDecoration(value: string): void;

  getFloor(): string;
  setFloor(value: string): void;

  getHouseYear(): string;
  setHouseYear(value: string): void;

  getStructBuild(): string;
  setStructBuild(value: string): void;

  getImage(): string;
  setImage(value: string): void;

  getFollow(): number;
  setFollow(value: number): void;

  getReleaseTime(): string;
  setReleaseTime(value: string): void;

  getTags(): number;
  setTags(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HouseSummary.AsObject;
  static toObject(includeInstance: boolean, msg: HouseSummary): HouseSummary.AsObject;
  static serializeBinaryToWriter(message: HouseSummary, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HouseSummary;
  static deserializeBinaryFromReader(message: HouseSummary, reader: jspb.BinaryReader): HouseSummary;
}

export namespace HouseSummary {
  export type AsObject = {
    id: number,
    houseNo: string,
    url: string,
    title: string,
    totalPrice: string,
    unitPrice: string,
    plot: string,
    region: string,
    layout: string,
    area: string,
    face: string,
    decoration: string,
    floor: string,
    houseYear: string,
    structBuild: string,
    image: string,
    follow: number,
    releaseTime: string,
    tags: number,
  }
}

export class House extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getHouseNo(): string;
  setHouseNo(value: string): void;

  getHouseInfo(): HouseInfo | undefined;
  setHouseInfo(value?: HouseInfo): void;
  hasHouseInfo(): boolean;
  clearHouseInfo(): void;

  getHouseBaseInfo(): HouseBaseInfo | undefined;
  setHouseBaseInfo(value?: HouseBaseInfo): void;
  hasHouseBaseInfo(): boolean;
  clearHouseBaseInfo(): void;

  getHouseTransactionInfo(): HouseTransactionInfo | undefined;
  setHouseTransactionInfo(value?: HouseTransactionInfo): void;
  hasHouseTransactionInfo(): boolean;
  clearHouseTransactionInfo(): void;

  getHousePicsList(): Array<HousePic>;
  setHousePicsList(value: Array<HousePic>): void;
  clearHousePicsList(): void;
  addHousePics(value?: HousePic, index?: number): HousePic;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): House.AsObject;
  static toObject(includeInstance: boolean, msg: House): House.AsObject;
  static serializeBinaryToWriter(message: House, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): House;
  static deserializeBinaryFromReader(message: House, reader: jspb.BinaryReader): House;
}

export namespace House {
  export type AsObject = {
    id: string,
    houseNo: string,
    houseInfo?: HouseInfo.AsObject,
    houseBaseInfo?: HouseBaseInfo.AsObject,
    houseTransactionInfo?: HouseTransactionInfo.AsObject,
    housePicsList: Array<HousePic.AsObject>,
  }
}

export class HouseInfo extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getHouseNo(): string;
  setHouseNo(value: string): void;

  getUrl(): string;
  setUrl(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  getSubTitle(): string;
  setSubTitle(value: string): void;

  getRegion(): string;
  setRegion(value: string): void;

  getTotalPrice(): string;
  setTotalPrice(value: string): void;

  getUnitPrice(): string;
  setUnitPrice(value: string): void;

  getRoomInfo(): string;
  setRoomInfo(value: string): void;

  getRoomSub(): string;
  setRoomSub(value: string): void;

  getTypeInfo(): string;
  setTypeInfo(value: string): void;

  getTypeSub(): string;
  setTypeSub(value: string): void;

  getAreaInfo(): string;
  setAreaInfo(value: string): void;

  getAreaSub(): string;
  setAreaSub(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HouseInfo.AsObject;
  static toObject(includeInstance: boolean, msg: HouseInfo): HouseInfo.AsObject;
  static serializeBinaryToWriter(message: HouseInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HouseInfo;
  static deserializeBinaryFromReader(message: HouseInfo, reader: jspb.BinaryReader): HouseInfo;
}

export namespace HouseInfo {
  export type AsObject = {
    id: string,
    houseNo: string,
    url: string,
    title: string,
    subTitle: string,
    region: string,
    totalPrice: string,
    unitPrice: string,
    roomInfo: string,
    roomSub: string,
    typeInfo: string,
    typeSub: string,
    areaInfo: string,
    areaSub: string,
  }
}

export class HouseBaseInfo extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getHouseNo(): string;
  setHouseNo(value: string): void;

  getLayout(): string;
  setLayout(value: string): void;

  getFloor(): string;
  setFloor(value: string): void;

  getAreaBuild(): string;
  setAreaBuild(value: string): void;

  getStructHouse(): string;
  setStructHouse(value: string): void;

  getAreaInner(): string;
  setAreaInner(value: string): void;

  getBuildType(): string;
  setBuildType(value: string): void;

  getFace(): string;
  setFace(value: string): void;

  getStructBuild(): string;
  setStructBuild(value: string): void;

  getDecoration(): string;
  setDecoration(value: string): void;

  getElevatorRatio(): string;
  setElevatorRatio(value: string): void;

  getElevator(): string;
  setElevator(value: string): void;

  getProperty(): string;
  setProperty(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HouseBaseInfo.AsObject;
  static toObject(includeInstance: boolean, msg: HouseBaseInfo): HouseBaseInfo.AsObject;
  static serializeBinaryToWriter(message: HouseBaseInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HouseBaseInfo;
  static deserializeBinaryFromReader(message: HouseBaseInfo, reader: jspb.BinaryReader): HouseBaseInfo;
}

export namespace HouseBaseInfo {
  export type AsObject = {
    id: string,
    houseNo: string,
    layout: string,
    floor: string,
    areaBuild: string,
    structHouse: string,
    areaInner: string,
    buildType: string,
    face: string,
    structBuild: string,
    decoration: string,
    elevatorRatio: string,
    elevator: string,
    property: string,
  }
}

export class HouseTransactionInfo extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getHouseNo(): string;
  setHouseNo(value: string): void;

  getListingTime(): string;
  setListingTime(value: string): void;

  getTradingAuthority(): string;
  setTradingAuthority(value: string): void;

  getLastTransaction(): string;
  setLastTransaction(value: string): void;

  getHousingPurposes(): string;
  setHousingPurposes(value: string): void;

  getHouseYear(): string;
  setHouseYear(value: string): void;

  getPropertyRights(): string;
  setPropertyRights(value: string): void;

  getMortgageInfo(): string;
  setMortgageInfo(value: string): void;

  getDocumentPhoto(): string;
  setDocumentPhoto(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HouseTransactionInfo.AsObject;
  static toObject(includeInstance: boolean, msg: HouseTransactionInfo): HouseTransactionInfo.AsObject;
  static serializeBinaryToWriter(message: HouseTransactionInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HouseTransactionInfo;
  static deserializeBinaryFromReader(message: HouseTransactionInfo, reader: jspb.BinaryReader): HouseTransactionInfo;
}

export namespace HouseTransactionInfo {
  export type AsObject = {
    id: string,
    houseNo: string,
    listingTime: string,
    tradingAuthority: string,
    lastTransaction: string,
    housingPurposes: string,
    houseYear: string,
    propertyRights: string,
    mortgageInfo: string,
    documentPhoto: string,
  }
}

export class HousePic extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getHouseNo(): string;
  setHouseNo(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getPicSmallUrl(): string;
  setPicSmallUrl(value: string): void;

  getPicNormalUrl(): string;
  setPicNormalUrl(value: string): void;

  getPicLargeUrl(): string;
  setPicLargeUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HousePic.AsObject;
  static toObject(includeInstance: boolean, msg: HousePic): HousePic.AsObject;
  static serializeBinaryToWriter(message: HousePic, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HousePic;
  static deserializeBinaryFromReader(message: HousePic, reader: jspb.BinaryReader): HousePic;
}

export namespace HousePic {
  export type AsObject = {
    id: string,
    houseNo: string,
    description: string,
    picSmallUrl: string,
    picNormalUrl: string,
    picLargeUrl: string,
  }
}

